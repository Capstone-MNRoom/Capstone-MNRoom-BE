package helper

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentChargeResponse struct {
	TransactionID     string
	OrderID           string
	PaymentType       string
	BankTransfer      string
	GrossAmount       int
	VANumber          string
	TransactionStatus string
}

type MidtransCustomer struct {
	Username string
	Email    string
	Phone    string
	Address  string
}

type PaymentGateway interface {
	ChargeBankTransfer(orderID string, amount int, bank string, customer MidtransCustomer) (*PaymentChargeResponse, error)
	VerifySignature(orderID, statusCode, grossAmount, signatureKey string) bool
	NormalizeStatus(transactionStatus, fraudStatus string) string
}

type midtransGateway struct {
	client    coreapi.Client
	serverKey string
}

func NewMidtransGateway() PaymentGateway {
	serverKey := os.Getenv("SERVER_KEY")
	var client coreapi.Client
	env := midtrans.Sandbox
	if strings.ToLower(os.Getenv("MIDTRANS_ENV")) == "production" {
		env = midtrans.Production
	}
	client.New(serverKey, env)
	return &midtransGateway{
		client:    client,
		serverKey: serverKey,
	}
}

func (m *midtransGateway) ChargeBankTransfer(orderID string, amount int, bank string, customer MidtransCustomer) (*PaymentChargeResponse, error) {
	bank = strings.ToLower(strings.TrimSpace(bank))

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(bank),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: customer.Username,
			Phone: customer.Phone,
			Email: customer.Email,
			BillAddr: &midtrans.CustomerAddress{
				Address: customer.Address,
			},
		},
	}

	res, err := m.client.ChargeTransaction(chargeReq)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response from payment gateway")
	}

	vaNumber := ""
	for _, v := range res.VaNumbers {
		vaNumber = v.VANumber
		break
	}
	// Fallback for Permata VA
	if vaNumber == "" && res.PermataVaNumber != "" {
		vaNumber = res.PermataVaNumber
	}

	return &PaymentChargeResponse{
		TransactionID:     res.TransactionID,
		OrderID:           res.OrderID,
		PaymentType:       res.PaymentType,
		BankTransfer:      bank,
		GrossAmount:       amount,
		VANumber:          vaNumber,
		TransactionStatus: res.TransactionStatus,
	}, nil
}

func (m *midtransGateway) VerifySignature(orderID, statusCode, grossAmount, signatureKey string) bool {
	raw := fmt.Sprintf("%s%s%s%s", orderID, statusCode, grossAmount, m.serverKey)
	hash := sha512.Sum512([]byte(raw))
	expected := hex.EncodeToString(hash[:])
	return expected == signatureKey
}

func (m *midtransGateway) NormalizeStatus(transactionStatus, fraudStatus string) string {
	switch transactionStatus {
	case "capture":
		if fraudStatus == "challenge" {
			return "challenge"
		}
		return "success"
	case "settlement":
		return "success"
	case "pending":
		return "pending"
	case "deny", "cancel", "expire":
		return "failed"
	default:
		return transactionStatus
	}
}
