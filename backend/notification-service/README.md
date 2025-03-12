## **4️⃣ Notification Service (lostfound-notification-service)**

# Lost & Found - Notification Service

## Overview
Sends SMS, email, and push notifications to users.

## Features
- Email notifications (SMTP)
- SMS alerts (Twilio)
- Push notifications (Firebase)

## Tech Stack
- Language: Python
- Messaging: Kafka / AWS SNS
- API: REST

## Installation
```bash
git clone https://github.com/your-org/lostfound-notification-service.git
cd lostfound-notification-service
pip install -r requirements.txt
python app.py
API Endpoints
```

## Method	Endpoint	Description
- POST	/notify/email	Send email alert
- POST	/notify/sms	Send SMS alert
- POST	/notify/push	Send push alert