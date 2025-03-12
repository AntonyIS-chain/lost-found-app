
## **2️⃣ Document Service (lostfound-document-service)**

# Lost & Found - Document Service

## Overview
Handles lost & found document reporting and matching.

## Features
- Users report lost or found documents
- AI-based document matching
- Image storage (AWS S3)

## Tech Stack
- Language: Golang
- Database: PostgreSQL
- Storage: AWS S3
- API: REST + GraphQL

## Installation
```bash
git clone https://github.com/your-org/lostfound-document-service.git
cd lostfound-document-service
go mod tidy
go run main.go
```

### API Endpoints
- Method	Endpoint	Description
- POST	/documents/report	Report lost or found doc
- GET	/documents/:id	Get document details
- GET	/documents/match	Find matching documents
