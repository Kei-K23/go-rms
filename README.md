# Go Restaurant Management System backend API

This is a general restaurant management system backend API with including JWT authentication and Stripe payment implemented in Go.

## Teach stack

- Go + Fiber
- Stripe
- MySQL
- more go tools...

## Features

- User register and login
- CRUD for restaurants, tables, categories, menus and orders
- JWT authentication
- Stripe payment support

## Prerequisites

- Go programming language installed on your system
- MySQL or another compatible database installed and running
- `.env` file containing environment variables (e.g., database connection details, server port)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/Kei-K23/go-rms.git
```

2. Navigate to the project directory:

```bash
cd go-rms

```

3. Create a .env file in the project root and add the following environment variables:

```bash
SECRET_KEY=<YOUR_SECRET_KEY>
PRODUCTION_DB_CONNECTION_STRING=<YOUR_PRODUCTION_DB_CONNECTION_STRING>
STRIPE_API_KEY=<YOUR_STRIPE_API_KEY>
```

## Usage

1. Install dependencies:

```bash
go mod tidy
```

1. Run migration

```bash
make migration
```

2. Push database table

```bash
make migrate-up
```

2. Run server

```bash
make run
```

This will serve the server at `http://localhost:4000`

1. Access the API endpoints using tools like cURL or Postman.

## API endpoints

All endpoints are available under `http://localhost:4000/api/v1`. Make sure prefix with your localhost with `/api/v1`. All these endpoints are test with Postman.

Authentication

- `POST /register`: Register a new user
- `POST /login`: Login user and get back access token to use at JWT bearer authentication

### Below defined endpoints are protected with JWT authentication. Make sure valid JWT token exist in bearer authentication header

Users

- `GET /users`: Get user according JWT token
- `PUT /users`: Update user
- `DELETE /users`: Delete the user

Restaurants

- `POST /restaurants`: Create new restaurant
- `GET /restaurants/:id`: Retrieve specific restaurant by ID
- `PUT /restaurants/:id`: Update restaurant by ID
- `DELETE /restaurants/:id`: Delete restaurant by ID

Tables

- `POST /restaurants/:id/tables`: Create new table for restaurant
- `GET /restaurants/:id/tables/:tid`: Retrieve specific table by ID
- `PUT /restaurants/:id/tables/:tid`: Update table by ID
- `DELETE /restaurants/:id/tables/:tid`: Delete table by ID

Categories

- `POST /categories`: Create new category
- `GET /categories`: Retrieve all category
- `GET /categories/:id`: Retrieve specific category by ID
- `PUT /categories/:id`: Update category by ID
- `DELETE /categories/:id`: Delete category by ID

Menus

- `POST /restaurants/:restaurantId/menus`: Create new menu for restaurant
- `GET /restaurants/:restaurantId/menus`: Retrieve all menus for a restaurant
- `GET /restaurants/:restaurantId/menus/:menuId`: Retrieve specific menu by ID
- `PUT /restaurants/:restaurantId/menus/:menuId`: Update menu by ID
- `DELETE /restaurants/:restaurantId/menus/:menuId`: Delete menu by ID

Order Items

- `POST /restaurants/:restaurantId/orderitems`: Create new order item
- `PUT /restaurants/:restaurantId/orderitems/:orderItemId`: Update order item by ID

Order

- `POST /restaurants/:restaurantId/orders`: Create new order
- `PUT /restaurants/:restaurantId/orders/:orderId`: Update order by ID

Payment

- `POST /checkout`: Create new checkout session for Stripe
- `PUT /success`: Route to use Stipe after successfully create payment
- `PUT /cancel`: Route to use Stipe when payment cancel due to error

## Todo

- Test all api endpoints
- Add endpoint for staff
- Add templates
