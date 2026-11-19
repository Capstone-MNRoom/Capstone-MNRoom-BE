package helper

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeStatus(t *testing.T) {
	gw := &midtransGateway{}

	tests := []struct {
		name              string
		transactionStatus string
		fraudStatus       string
		expected          string
	}{
		{"capture with challenge", "capture", "challenge", "challenge"},
		{"capture with accept", "capture", "accept", "success"},
		{"settlement", "settlement", "", "success"},
		{"pending", "pending", "", "pending"},
		{"deny", "deny", "", "failed"},
		{"cancel", "cancel", "", "failed"},
		{"expire", "expire", "", "failed"},
		{"unknown status", "custom_status", "", "custom_status"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := gw.NormalizeStatus(tt.transactionStatus, tt.fraudStatus)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestVerifySignature(t *testing.T) {
	serverKey := "SB-Mid-server-TESTKEY123"
	gw := &midtransGateway{serverKey: serverKey}

	orderID := "ORDER-12345"
	statusCode := "200"
	grossAmount := "500000.00"

	raw := fmt.Sprintf("%s%s%s%s", orderID, statusCode, grossAmount, serverKey)
	hash := sha512.Sum512([]byte(raw))
	validSig := hex.EncodeToString(hash[:])

	t.Run("Valid Signature", func(t *testing.T) {
		assert.True(t, gw.VerifySignature(orderID, statusCode, grossAmount, validSig))
	})

	t.Run("Invalid Signature", func(t *testing.T) {
		assert.False(t, gw.VerifySignature(orderID, statusCode, grossAmount, "invalidsignaturehash123"))
	})
}
