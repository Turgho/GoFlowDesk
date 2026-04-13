package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = 64 * 1024
	iterations  = 1
	parallelism = 4
	keyLength   = 32
	saltLength  = 16
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, encoded string) bool
}

type argon2Hasher struct{}

func NewArgon2Hasher() PasswordHasher {
	return &argon2Hasher{}
}

func (h *argon2Hasher) HashPassword(password string) (string, error) {
	// Gera um salt aleatório 16 bytes
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Gera o hash usando Argon2id
	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Codifica o salt e o hash em base64 para armazenamento
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, iterations, parallelism, b64Salt, b64Hash,
	)

	// Retorna a string codificada
	return encoded, nil
}

func (h *argon2Hasher) ComparePassword(password, encoded string) bool {
	// Divide a string codificada em partes
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return false
	}

	var mem uint32
	var iter uint32
	var par uint8

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &iter, &par)
	if err != nil {
		return false
	}

	// Decodifica o salt e o hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	// Decodifica o hash armazenado
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// Gera um hash do password fornecido usando o mesmo salt e parâmetros
	newHash := argon2.IDKey([]byte(password), salt, iter, mem, par, uint32(len(hash)))

	// Compara os hashes usando uma comparação de tempo constante
	return subtle.ConstantTimeCompare(hash, newHash) == 1
}
