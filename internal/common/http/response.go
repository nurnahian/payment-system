package httphelper

import (
	"encoding/json"
	"net/http"
	"payment-system/internal/common/crypto"
)

func WriteEncryptedJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}

	encrypted, err := crypto.EncryptAES(string(jsonBytes))
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(encrypted))
}
