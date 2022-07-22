package midtrans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/midtrans/midtrans-go"
)

type Client struct {
	ServerKey  string
	ClientKey  string
	Env        midtrans.EnvironmentType
	HttpClient midtrans.HttpClient
	Options    *midtrans.ConfigOptions
}

func setupGlobalMidtransConfigApi() {
	midtrans.ServerKey = os.Getenv("SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
}

func (c Client) ChargeTransaction(req *ChargeReq) (*ChargeResponse, *midtrans.Error) {
	resp := &ChargeResponse{}
	jsonReq, _ := json.Marshal(req)
	err := c.HttpClient.Call(http.MethodPost,
		fmt.Sprintf("%s/v2/charge", c.Env.BaseUrl()),
		&c.ServerKey,
		c.Options,
		bytes.NewBuffer(jsonReq),
		resp,
	)

	if err != nil {
		return resp, err
	}
	return resp, nil
}
