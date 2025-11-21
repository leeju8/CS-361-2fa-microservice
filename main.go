package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net/http"

	"github.com/pquerna/otp/totp"
)

type GenerateResponse struct {
	Secret string `json:"secret"`
	QRCode string `json:"qr_code"` // base64 encoded PNG
	URL    string `json:"url"`
}

type VerifyRequest struct {
	Secret string `json:"secret"`
	Code   string `json:"code"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
}

func generate(w http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		http.Error(w, "Failed to generate key", http.StatusInternalServerError)
		log.Printf("Error generating key: %v", err)
		return
	}

	// Generate QR code
	img, err := key.Image(200, 200)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		log.Printf("Error generating QR code: %v", err)
		return
	}

	// Convert image to base64
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		http.Error(w, "Failed to encode QR code", http.StatusInternalServerError)
		log.Printf("Error encoding QR code: %v", err)
		return
	}
	qrCodeBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	response := GenerateResponse{
		Secret: key.Secret(),
		QRCode: qrCodeBase64,
		URL:    key.URL(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func verify(w http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var verifyReq VerifyRequest
	if err := json.NewDecoder(req.Body).Decode(&verifyReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Verify the code
	valid := totp.Validate(verifyReq.Code, verifyReq.Secret)

	response := VerifyResponse{
		Valid: valid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/verify", verify)

	fmt.Println("2FA Microservice running on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
