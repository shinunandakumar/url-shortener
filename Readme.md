# URL Shortener Project

This project is a URL shortener service that allows users to generate shortened URLs and redirect to the original URLs.

## APIs

### 1. Generate Shortened URL

- **Endpoint:** `api/v1/generate-url`
- **Description:** Generates a shortened URL for a given original URL.
- **Method:** POST
- **Request:**
  - **Body:**
    - `original_url` (string): The original URL to be shortened.
- **Response:**
  - `shortened_url` (string): The shortened URL.

### 2. Health Check

- **Endpoint:** `api/v1/health-check`
- **Description:** Provides a health check for the service.
- **Method:** GET
- **Response:**
  - `status` (string): A message indicating the health status of the service.

### 3. Metrics Data

- **Endpoint:** `api/v1/metrics`
- **Description:** Provides metrics data for monitoring and analysis.
- **Method:** GET
- **Response:**
  - Metrics data in the desired format (e.g., JSON, CSV, etc.).

### 4. Redirection API

- **Endpoint:** `/`
- **Description:** Handles redirection to the original URL using the shortened URL.
- **Method:** GET
- **Request:**
  - **URL Parameter:**
    - `shortened_url` (string): The shortened URL to redirect.
- **Response:**
  - Redirects to the original URL.

## Setup

### Prerequisites

- Go (version 1.16+)
- Redis (for storing URL mappings)

### Installation

1. Clone this repository: `git clone https://github.com/your-username/url-shortener.git`
2. Navigate to the project directory: `cd url-shortener`
3. Install dependencies: `go mod download`
4. Build the project: `go build -o url-shortener`

## Usage

1. Run the server: `./url-shortener`
2. Access the APIs using the provided endpoints and methods.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
