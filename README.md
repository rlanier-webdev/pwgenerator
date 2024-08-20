# Password Generator API

## Overview

This project provides a simple web server with an endpoint to generate random passwords. It uses the Gin framework for handling HTTP requests and serves both HTML and API responses.

## Features

- **Web Interface**: Displays a page with a randomly generated password.
- **API Endpoint**: Provides an endpoint to generate and return a password in JSON format.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- [Gin framework](https://github.com/gin-gonic/gin)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/rlanier-webdev/pwgenerator.git
   cd pwgenerator
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run main.go
   ```

   The server will start and listen on port `8080`.

### Endpoints

- **Web Interface**

  - **URL**: `http://localhost:8080/`
  - **Method**: `GET`
  - **Description**: Serves an HTML page with a randomly generated password.

- **API Endpoint**

  - **URL**: `http://localhost:8080/api/password`
  - **Method**: `GET`
  - **Description**: Returns a JSON object with a randomly generated password.

  **Response Example:**
  ```json
  {
      "password": "GeneratedPasswordHere"
  }
  ```

### Code Structure

- **`main.go`**: Contains the main application logic and routes. Includes password generation functions and API route definitions.
- **`handler.go`**: Defines the handler for the web interface.