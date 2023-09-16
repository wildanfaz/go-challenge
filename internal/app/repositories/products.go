package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/wildanfaz/go-challenge/internal/app/entities"
	"github.com/wildanfaz/go-challenge/pkg/utils"
)

type Products interface {
	AddProductToCart(ctx context.Context, product entities.AddProductToCart) error
	ListProducts(ctx context.Context, product entities.Product) (*entities.Products, error)
	CheckProduct(ctx context.Context, productID string) (bool, error)
	ListProductsInCart(ctx context.Context) (*entities.ProductsInCart, error)
	DeleteInShoppingCart(ctx context.Context, productID string) error
	CheckProductInCart(ctx context.Context, productID string) (bool, error)
	CheckBalance(ctx context.Context) (int, bool, error)
	CheckoutAndPayment(ctx context.Context, balance int) error
}

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) Products {
	return &ProductsRepo{db: db}
}

func (str *ProductsRepo) AddProductToCart(ctx context.Context, product entities.AddProductToCart) error {
	var (
		count int
	)

	q := `INSERT INTO carts(user_id, product_id, amount) VALUES(?,?,?)`

	if err := str.db.QueryRowContext(ctx, `SELECT count(p.id) 
	FROM products p
	JOIN carts c ON c.product_id=p.id
	LEFT JOIN payments pm ON pm.cart_id=c.id
	WHERE c.user_id=? AND c.product_id=? AND pm.cart_id IS NULL`, product.UserID, product.ProductID).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		if _, err := str.db.ExecContext(ctx, `UPDATE carts SET amount = amount + ? WHERE user_id=? AND product_id=?`,
			product.Amount, product.UserID, product.ProductID); err != nil {
			return err
		}
	} else {
		_, err := str.db.ExecContext(ctx, q, product.UserID, product.ProductID, product.Amount)

		if err != nil {
			return err
		}
	}

	return nil
}

func (str *ProductsRepo) ListProducts(ctx context.Context, product entities.Product) (*entities.Products, error) {
	var (
		scanProduct entities.Product
		products    entities.Products
	)

	keys, values := utils.StructToKeyValue(product, "json")

	q := `SELECT id, name, description, category, price, created_at, updated_at FROM products ORDER BY created_at DESC`

	if len(values) > 0 && len(keys) > 0 {
		q = strings.Replace(q, "products", "products WHERE %s", 1)

		var wq string
		for _, v := range keys {
			wq += v + "=" + "?" + ","
		}

		wq, _ = strings.CutSuffix(wq, ",")

		q = fmt.Sprintf(q, wq)
	}

	rows, err := str.db.QueryContext(ctx, q, values...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&scanProduct.ID, &scanProduct.Name, &scanProduct.Description,
			&scanProduct.Category, &scanProduct.Price, &scanProduct.CreatedAt, &scanProduct.UpdatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, scanProduct)
	}

	return &products, nil
}

func (str *ProductsRepo) CheckProduct(ctx context.Context, productID string) (bool, error) {
	var (
		count int
	)

	if err := str.db.QueryRowContext(ctx, "SELECT count(id) FROM products WHERE id=?", productID).
		Scan(&count); err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}

	return true, nil
}

func (str *ProductsRepo) ListProductsInCart(ctx context.Context) (*entities.ProductsInCart, error) {
	var (
		scanProduct entities.ProductInCart
		products    entities.ProductsInCart
		userID      = ctx.Value("user_id")
	)

	q := `SELECT p.id, p.name, p.description, p.category, p.price, c.amount, p.created_at, p.updated_at 
	FROM products p
	JOIN carts c ON c.product_id=p.id
	LEFT JOIN payments pm ON pm.cart_id=c.id
	WHERE c.user_id=? AND pm.cart_id IS NULL
	ORDER BY c.created_at DESC
	`

	rows, err := str.db.QueryContext(ctx, q, userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&scanProduct.ID, &scanProduct.Name, &scanProduct.Description,
			&scanProduct.Category, &scanProduct.Price, &scanProduct.Amount,
			&scanProduct.CreatedAt, &scanProduct.UpdatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, scanProduct)
	}

	return &products, nil
}

func (str *ProductsRepo) DeleteInShoppingCart(ctx context.Context, productID string) error {
	if _, err := str.db.ExecContext(ctx, "DELETE FROM carts WHERE user_id=? AND product_id=?",
		ctx.Value("user_id"), productID); err != nil {
		return err
	}

	return nil
}

func (str *ProductsRepo) CheckProductInCart(ctx context.Context, productID string) (bool, error) {
	var (
		count int
	)

	if err := str.db.QueryRowContext(ctx, "SELECT count(id) FROM carts WHERE user_id=? AND product_id=?",
		ctx.Value("user_id"), productID).
		Scan(&count); err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}

	return true, nil
}

func (str *ProductsRepo) CheckBalance(ctx context.Context) (int, bool, error) {
	var (
		balance int
		cost    int
	)

	q := `SELECT SUM(p.price * c.amount) as cost FROM products p 
	JOIN carts c ON c.product_id=p.id
	LEFT JOIN payments pm ON pm.cart_id=c.id
	WHERE c.user_id=? AND pm.cart_id IS NULL
	`

	if err := str.db.QueryRowContext(ctx, q, ctx.Value("user_id")).Scan(&cost); err != nil {
		return 0, false, err
	}

	if err := str.db.QueryRowContext(ctx, "SELECT balance FROM users WHERE id=?", ctx.Value("user_id")).Scan(&balance); err != nil {
		return 0, false, err
	}

	balance -= cost

	if balance < 0 {
		return 0, false, nil
	}

	return balance, true, nil
}

func (str *ProductsRepo) CheckoutAndPayment(ctx context.Context, balance int) error {
	var (
		cartID  string
		cartsID []string
	)

	q := `UPDATE users SET balance=? WHERE id=?`

	tx, err := str.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer tx.Commit()

	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, q, balance, ctx.Value("user_id")); err != nil {
		tx.Rollback()
		return err
	}

	rows, err := tx.QueryContext(ctx, `SELECT c.id 
	FROM carts c
	JOIN users u ON u.id=c.user_id
	LEFT JOIN payments p ON p.cart_id=c.id
	WHERE u.id=? AND p.cart_id IS NULL
	`, ctx.Value("user_id"))

	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {
		err = rows.Scan(&cartID)

		if err != nil {
			tx.Rollback()
			return err
		}

		cartsID = append(cartsID, cartID)
	}

	if len(cartsID) < 1 {
		return errors.New("carts not found")
	}

	for _, v := range cartsID {
		if _, err = tx.ExecContext(ctx, "INSERT INTO payments(cart_id) VALUES(?)", v); err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
