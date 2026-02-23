# #URL Shortener

## URL shortener developed in Go using Gin and Redis.

### Technologies

Go

Gin (HTTP framework)

Redis (storage for URL mappings)

### How to Run
Prerequisites

Go 1.21+

Redis running locally or via Docker

### Installation
git clone https://github.com/glrmrissi/url-shortener
cd url-shortener

cp .env.example .env

go mod tidy

go run main.go
Endpoints
Shorten URL
POST /shorten
Content-Type: application/json

{
  "url": "https://mysite.com/very-long-page"
}

Response:

{
  "code": "x7k2p1",
  "original_url": "https://mysite.com/very-long-page",
  "short_url": "http://localhost:8080/x7k2p1"
}
Redirect
GET /:code

Automatically redirects to the original URL. URLs expire after 30 days.