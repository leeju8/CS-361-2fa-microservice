# 2FA Microservice

A simple microservice that generates TOTP secrets and verifies 2FA codes.

# Running the Service

```bash
go mod tidy
go run main.go
```

The service will run on `http://localhost:8090`

# Requesting Data

### Generate 2FA Secret

**Endpoint:** `POST /generate`

**Request:** Client sends a POST request with an empty body

**Example:**

```bash
curl -X POST http://localhost:8090/generate
```

**Response (Success):** Client receives 200 OK Status Code and JSON body containing:

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

**Response (Failure):** Client receives 400 BAD REQUEST Status Code on invalid method request and 500 INTERNAL SERVER ERROR on server failure

### Verify 2FA Code

**Endpoint:** `POST /verify`

**Request:** Client sends a POST request with JSON body containing:

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

**Response (Success):** Client receives 200 OK Status Code and JSON body containing:

```json
{
  "valid": true
}
```

- `valid`: `true` if the code is correct, `false` otherwise

**Response (Failure):** Client receives 400 BAD REQUEST Status Code on invalid method request and 500 INTERNAL SERVER ERROR on server failure

### UML Sequence Diagram

<img width="420" height="440" alt="image" src="https://github.com/user-attachments/assets/ea0ee4db-ef66-49c0-b263-45a93bbc79fa" />
