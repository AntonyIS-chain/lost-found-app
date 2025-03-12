## **3️⃣ Payment Service (lostfound-payment-service)**

# Lost & Found - Payment Service

## Overview
Manages payments when a lost document is matched and recovered.

## Features
- Stripe, PayPal, M-Pesa integration
- Transaction history
- Secure payments processing

## Tech Stack
- Language: Golang
- Database: PostgreSQL
- Payment Gateway: Stripe, PayPal, M-Pesa
- API: REST + gRPC

## Installation
```bash
git clone https://github.com/your-org/lostfound-payment-service.git
cd lostfound-payment-service
go mod tidy
go run main.go
```

### API Endpoints
- Method	Endpoint	Description
- POST	/payments/pay	Make a payment
- GET	/payments/:id	Get transaction info



