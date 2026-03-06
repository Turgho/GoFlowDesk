package render

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Evita que clientes enviem corpos de requisição muito grandes,
// o que pode levar a ataques de negação de serviço (DoS)
// ou esgotamento de recursos do servidor.
const DefaultMaxBodySize = 1 << 20 // 1 MB

func WriteJSON(w http.ResponseWriter, status int, payload any, headers http.Header) error {
	// Adiciona os headers personalizados à resposta
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	return enc.Encode(payload)
}

func ReadJSON[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var body T

	if r.Body == nil {
		return body, fmt.Errorf("request body is empty")
	}

	// Limite padrão
	r.Body = http.MaxBytesReader(w, r.Body, DefaultMaxBodySize)

	contentType := r.Header.Get("Content-Type")
	if contentType != "" && !strings.HasPrefix(contentType, "application/json") {
		return body, fmt.Errorf("content type must be application/json")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		return body, err
	}

	if decoder.More() {
		return body, fmt.Errorf("request body must contain a single JSON object")
	}

	return body, nil
}
