# go-challenge

Skill test using fiber's framework

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Documentations](#documentations)

## Features

- Authentication
- Products's service

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/wildanfaz/go-challenge.git
   ```

2. Change to the project directory:

   ```sh
   cd go-challenge
   ```

3. Run the project:

   ```sh
   docker compose -f ./deployments/docker-compose.yml up
   ```

## Usage

Install dependencies
  ```sh
   make install
   ```

Start app
  ```sh
   make start
   ```

Migrate database
  ```sh
   make migrate
   ```

Rollback database
  ```sh
  make rollback
   ```

Add user's balance (change email in the Makefile if needed)
  ```sh
   make add_balance
   ```

Add products's dumy
 ```sh
   make dumy
   ```

## Documentations

Postman

https://documenter.getpostman.com/view/22978251/2s9YC7SXAG

ERD

https://dbdiagram.io/d/go-challenge-6506a10c02bd1c4a5eb6fc03