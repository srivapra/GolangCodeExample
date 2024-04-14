# Go URL Shortener Service

This is a simple URL shortener service implemented in Go. It provides a RESTful API for shortening URLs, redirecting users to original URLs, and fetching metrics about the most frequently shortened domains.

# Setup
1. Clone the repository: 
    - git clone "repository-url"
2. Navigate to the project directory: 
    - cd "project-directory"

# Running the Server
1. Build the application:
    - go build
2. Start the server: 
    - ./url-shortener

The server will run on http://localhost:8080 by default.

# Endpoints
- Shorten URL: 'POST /shorten'
    - Request Body: { "url": "original-url-here" }
    - Response Body: { "shortenedUrl": "shortened-url-here" }

- Redirection: 'GET /shortened-url-here'
    - Redirects to the original URL.

- Metrics: 'GET /metrics'
    - Response Body: 
    [ {"domain":"domain-name","count":count},{"domain":"domain-name","count":count},{"domain":"domain-name","count":count} ]

# Dockerization
1. Build Docker image :
    - docker build -t url-shortener .

2. Run Docker container :
    - docker run -d -p 8080:8080 --name url-shortener-app url-shortener

# Accessing the Services
- The service will be available at http://localhost:8080
- Access metrics API: http://localhost:8080/metrics