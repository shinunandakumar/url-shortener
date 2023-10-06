# URL Shortener Project

This project is a URL shortener service that allows users to generate shortened URLs and redirect to the original URLs.

## APIs

### 1. Generate Shortened URL

- **Endpoint:** `api/v1/generate-url`
- **Description:** Generates a shortened URL for a given original URL.
- **Method:** POST
- **Request:**
  - **Body:**
    - `url` (string): The original URL to be shortened.
- **Response:**
  - `url` (string): The shortened URL.
  - `status_code` (int): Status code.
  

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
  - 'Metrics data as a list format'

### 4. Redirection API

- **Endpoint:** `/`
- **Description:** Handles redirection to the original URL using the shortened URL.
- **Method:** GET
- **Request:**
  - **URL Parameter:**
    - `url` (string): The shortened URL to redirect.
- **Response:**
  - Redirects to the original URL.

## Setup

### Prerequisites

- docker 
- docker-compose or docker compose plugin

### Installation

1. Clone this repository: `git clone https://github.com/your-username/url-shortener.git`
2. Navigate to the project directory: `cd url-shortener`
3. Install dependencies: `https://docs.docker.com/engine/install/`
4. Build the project: `make docker-build`

## Usage

1. Run the server: `make docker-run`
2. Access the APIs using the provided endpoints and methods.
3. Run `make help` for further commands to use

