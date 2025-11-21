# 2FA Microservice

A simple microservice that generates TOTP secrets and verifies 2FA codes.

# Running the Service

go mod tidy

go run main.go

The service will run on `http://localhost:8090`

# Requesting Data

### Generate 2FA Secret

**Endpoint:** `POST /generate`

**Request:** Send a POST request with an empty body

**Example:**

```bash
curl -X POST http://localhost:8090/generate
```

### Verify 2FA Code

**Endpoint:** `POST /verify`

**Request:** Send a POST request with JSON body containing:

- `secret`: The TOTP secret (returned from /generate)
- `code`: The 6-digit code from authenticator app

**Example:**

```bash
curl -X POST http://localhost:8090/verify \
  -H "Content-Type: application/json" \
  -d '{
    "secret": "SECRET_VALUE",
    "code": "123456"
  }'
```

# Receiving Data

### Generate Response

```json
{
  "secret": "JBSWY3DPEHPK3PXP",
  "qr_code": "base64_encoded_png_image",
  "url": "otpauth://totp/MyApp:user@example.com?secret=JBSWY3DPEHPK3PXP&issuer=MyApp"
}
```

- `secret`: Secret needed for verification. Store this value
- `qr_code`: Base64 encoded PNG image to scan with authenticator app
- `url`: The otpauth URL (can be used instead of QR code)

### Verify Response

```json
{
  "valid": true
}
```

- `valid`: `true` if the code is correct, `false` otherwise

### UML Sequence Diagram

<img width="376" height="330" alt="image" src="https://github.com/user-attachments/assets/4c5b6df9-95ad-4524-b387-0c1998700d88" />
