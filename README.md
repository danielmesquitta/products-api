# Products API

This is an Products API CRUDproducts.
It is written in Go and uses the Echo framework and a MySQL database.

## Dependencies

1. [Go 1.22.4](https://go.dev/doc/install)
2. [Make](https://www.gnu.org/software/make/)
3. [Docker](https://www.docker.com/) with a MySQL container or [MySQL](https://www.mysql.com/downloads/) itself

## Running the API

1. Clone the repository.
2. Install the dependencies with `make install`.
3. Copy `.env.example` to `.env` and fill in your environment variables.
4. Run the migrations with `make migrations_up`
5. Finally, start the application with live-reload enabled running `make dev`.

## API Endpoints

All API endpoints are already documented in the Swagger available at the `/api/v1/docs/index.html` route.
Here is an overview:

- `GET /api/v1/products`: List all products.
- `POST /api/v1/products`: Create a new product.
- `GET /api/v1/products/{id}`: Get a specific product by its ID.
- `PUT /api/v1/products/{id}`: Update a specific product by its ID.
- `DELETE /api/v1/products/{id}`: Delete a specific product by its ID.

## Testing

Unit tests can be run with `make test`.
