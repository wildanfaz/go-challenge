package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wildanfaz/go-challenge/internal/app/routers"
	"github.com/wildanfaz/go-challenge/migrations"
)

var email string

var root = &cobra.Command{
	Short: "Default Command",
}

var startApp = &cobra.Command{
	Use:   "start",
	Short: "Start Application",
	Run: func(cmd *cobra.Command, args []string) {
		routers.New()
	},
}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Table",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
	},
}

var rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback Table",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Rollback()
	},
}

var addBalance = &cobra.Command{
	Use:   "add_balance",
	Short: "Increase the User's Balance by the Amount of 1 million",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.AddBalance(email)
	},
}

var dumy = &cobra.Command{
	Use:   "dumy",
	Short: "Insert Products's Dumy",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Dumy()
	},
}

func Start() error {
	root.PersistentFlags().StringVar(&email, "email", "example@mail.com", "user's email")
	viper.BindPFlag("email", root.PersistentFlags().Lookup("email"))

	root.AddCommand(startApp)
	root.AddCommand(migrate)
	root.AddCommand(rollback)
	root.AddCommand(addBalance)
	root.AddCommand(dumy)

	return root.Execute()
}
