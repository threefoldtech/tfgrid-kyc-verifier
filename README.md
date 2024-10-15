# TFGrid KYC Service

## Overview

TFGrid KYC Service is a Go-based service that provides Know Your Customer (KYC) functionality for the TFGrid. It integrates with iDenfy for identity verification.

## Features

- Identity verification using iDenfy
- Blockchain integration with TFChain (Substrate-based)
- MongoDB for data persistence
- RESTful API endpoints for KYC operations
- Swagger documentation
- Containerized deployment

## Prerequisites

- Go 1.22+
- MongoDB 4.4+
- Docker and Docker Compose (for containerized deployment)
- iDenfy API credentials

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/tfgrid-kyc-verifier.git
    cd tfgrid-kyc-verifier
    ```

2. Set up your environment variables:

    ```bash
    cp .app.env.example .app.env
    cp .db.env.example .db.env
    ```

Edit `.app.env` and `.db.env` with your specific configuration details.

## Configuration

The application uses environment variables for configuration. Here's a list of all available configuration options:

### Database Configuration

- `MONGO_URI`: MongoDB connection URI (default: "mongodb://localhost:27017")
- `DATABASE_NAME`: Name of the MongoDB database (default: "tfgrid-kyc-db")

### Server Configuration

- `PORT`: Port on which the server will run (default: "8080")

### iDenfy Configuration

- `IDENFY_API_KEY`: API key for iDenfy service
- `IDENFY_API_SECRET`: API secret for iDenfy service
- `IDENFY_BASE_URL`: Base URL for iDenfy API (default: "<https://ivs.idenfy.com/api/v2>")
- `IDENFY_CALLBACK_SIGN_KEY`: Callback signing key for iDenfy webhooks
- `IDENFY_WHITELISTED_IPS`: Comma-separated list of whitelisted IPs for iDenfy callbacks

### TFChain Configuration

- `TFCHAIN_WS_PROVIDER_URL`: WebSocket provider URL for TFChain (default: "wss://tfchain.grid.tf")

### Rate Limiting

- `MAX_TOKEN_REQUESTS_PER_MINUTE`: Maximum number of token requests allowed per minute for same IP address (default: 4)

### Verification Settings

- `SUSPICIOUS_VERIFICATION_OUTCOME`: Outcome for suspicious verifications (default: "verified")
- `EXPIRED_DOCUMENT_OUTCOME`: Outcome for expired documents (default: "unverified")
- `CHALLENGE_WINDOW`: Time window (in seconds) for challenge validation (default: 120)
- `MIN_BALANCE_TO_VERIFY_ACCOUNT`: Minimum balance (in units TFT) required to verify an account (default: 10000000)

To configure these options, you can either set them as environment variables or include them in your `.env` file.

Refer to `internal/configs/config.go` for a full list of configuration options.

## Running the Application

### Using Docker Compose

To start the server and MongoDB using Docker Compose:

```bash
docker-compose up -d --build
```

### Running Locally

To run the application locally:

1. Ensure MongoDB is running and accessible.
2. export the environment variables:

    ```bash
    set -a
    source .app.env
    set +a
    ```

3. Run the application:

    ```bash
    go run cmd/api/main.go
    ```

## API Endpoints

### Client endpoints

- `POST /api/v1/token`: Get or create a verification token
- `GET /api/v1/data`: Get verification data
- `GET /api/v1/status`: Get verification status

### Webhook endpoints

- `POST /webhooks/idenfy/verification-update`: Process verification update (webhook)
- `POST /webhooks/idenfy/id-expiration`: Process document expiration notification (webhook)

Refer to the Swagger documentation for detailed information on request/response formats.

## Swagger Documentation

Swagger documentation is available. To view it, run the application and navigate to the `/docs` endpoint in your browser.

## Project Structure

- `cmd/`: Application entrypoints
- `internal/`: Internal packages
  - `configs/`: Configuration handling
  - `handlers/`: HTTP request handlers
  - `models/`: Data models
  - `responses/`: API response structures
  - `services/`: Business logic
  - `repositories/`: Data repositories
  - `middlewares/`: Middlewares
  - `clients/`: External clients
  - `server/`: Server and router setup
- `api/docs/`: Swagger documentation
- `scripts/`: Development and utility scripts
- `docs/`: Documentation

## Development

### Running Tests

To run the test suite:

TODO: Add tests

### Building the Docker Image

To build the Docker image:

```bash
docker build -t tfgrid-kyc-service .
```

### Running the Docker Container

To run the Docker container and use .env variables:

```bash
docker run -d -p 8080:8080 --env-file .app.env tfgrid-kyc-service
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache 2.0 License. See the `LICENSE` file for more details.
