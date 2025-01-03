E-Commerce REST API

Welcome to the E-Commerce REST API! This project provides a robust backend system for managing an e-commerce platform, including user management, product management, and order processing. The API is built with Go (Golang) for high performance and scalability.

🌟 API Documentation

Explore the complete API documentation for all endpoints and usage details:

👉 View the API Documentation Here

The documentation includes:

A detailed list of all API endpoints.

Request and response examples.

Error handling guidelines.

Authentication and authorization requirements.

Features

User Authentication: Secure user registration, login, and JWT-based authentication.

Product Management: Add, update, delete, and fetch product details.

Order Management: Place, update, and track orders.

Admin Panel: Manage users, products, and orders with admin privileges.

RESTful Endpoints: Well-structured API following REST principles.

Installation and Setup

Prerequisites

Go (version 1.19 or later)

A database system (e.g., PostgreSQL or MongoDB)

Git for version control

Steps

Clone the repository:

git clone https://github.com/Nyxoy77/REST_API_GO.git
cd REST_API_GO

Install dependencies:

go mod tidy

Configure environment variables:

Create a .env file in the root directory.

Add the following variables:

DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=ecommerce
JWT_SECRET=your_jwt_secret

Run database migrations (if applicable).

Start the server:

go run main.go

Access the API at http://localhost:8080.

API Endpoints (Overview)

User Routes

POST /api/v1/register - Register a new user.

POST /api/v1/login - Authenticate and retrieve a JWT token.

GET /api/v1/user/:id - Fetch user details (requires authentication).

Product Routes

POST /api/v1/products - Add a new product (Admin only).

GET /api/v1/products - Fetch all products.

GET /api/v1/products/:id - Fetch details of a single product.

PUT /api/v1/products/:id - Update a product (Admin only).

DELETE /api/v1/products/:id - Delete a product (Admin only).

Order Routes

POST /api/v1/orders - Place a new order.

GET /api/v1/orders/:id - Fetch order details.

PUT /api/v1/orders/:id - Update order status (Admin only).

For detailed examples and payload structures, check the API Documentation.

Tech Stack

Programming Language: Go (Golang)

Database: PostgreSQL or MongoDB (configurable)

Authentication: JWT-based secure authentication

API Framework: net/http with middleware for routing and authentication

Testing

Run unit tests using:

go test ./...

Deployment

Use Docker to containerize and deploy the application.

Example Docker commands:

docker build -t ecommerce-api .
docker run -p 8080:8080 ecommerce-api

Contributing

We welcome contributions! Follow these steps:

Fork the repository.

Create a feature branch:

git checkout -b feature/your-feature

Commit your changes:

git commit -m "Add your message here"

Push to the branch:

git push origin feature/your-feature

Open a pull request.

License

This project is licensed under the MIT License.

Contact

Author: Shivam (GitHub: Nyxoy77)

API Documentation: View Here

Thank you for using the E-Commerce REST API! Feel free to reach out for any questions or suggestions.
