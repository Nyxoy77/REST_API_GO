# E-Commerce REST API

**E-Commerce REST API** provides a backend system for managing users, products, and orders using **Go (Golang)**.

---

## API Documentation

👉 [API Documentation](https://nyxoy77.github.io/REST_API_GO/)

Includes:
- API endpoints
- Request/response examples
- Authentication requirements

---

## Features

- Secure user authentication with **JWT**
- **CRUD operations** for products and orders
- Admin privileges for management
- RESTful design principles

---

## Installation

### Prerequisites

- **Go (1.19+)**
- **Database** (PostgreSQL/MongoDB)
- **Git**

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/Nyxoy77/REST_API_GO.git
   cd REST_API_GO
   
### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/Nyxoy77/REST_API_GO.git
   cd REST_API_GO
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure `.env`:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=username
   DB_PASSWORD=password
   DB_NAME=ecommerce
   JWT_SECRET=your_jwt_secret
   ```

4. Start the server:
   ```bash
   go run main.go
   ```

5. API available at `http://localhost:8080`.

---

## API Endpoints

### Public Routes

- `POST /login` - User login
- `POST /register` - User registration
- `POST /forgot` - Request password reset
- `PUT /update_pass/{token}` - Update password using token
- `POST /refresh` - Refresh JWT token

### Protected Routes (Authenticated)

- `GET /products` - Fetch all products

### Admin Routes

- `GET /get_all_users` - Fetch all users
- `GET /get_admins` - Fetch all admins
- `GET /get_customers` - Fetch all customers
- `POST /add_product` - Add a new product
- `DELETE /remove_product` - Remove a product
- `PUT /update_price` - Update product price

### User Routes

- `POST /add_to_cart` - Add item to cart
- `DELETE /remove_item` - Remove item from cart
- `POST /order` - Place an order

---

## Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL/MongoDB/Supabase
- **Authentication**: JWT
- **Framework**: `gorilla/mux`

---

## Testing

Run tests:
```bash
go test ./...
```

---

## Deployment

Use Docker:
```bash
docker build -t ecommerce-api .
docker run -p 8080:8080 ecommerce-api
```

---

## Contributing

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit changes:
   ```bash
   git commit -m "Your message"
   ```
4. Push to branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request.

---

## License

Licensed under the [MIT License](LICENSE).

---

## Contact

- **Author**: [Nyxoy77](https://github.com/Nyxoy77)
- **API Documentation**: [Link](https://nyxoy77.github.io/REST_API_GO/)
```
