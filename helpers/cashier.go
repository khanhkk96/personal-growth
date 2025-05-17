package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"personal-growth/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

type MoMoPayload struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
	Lang        string `json:"lang"`
}

func PayViaMoMo(amountV int64, description string) (string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	a, _ := flake.NextID()
	b, err := flake.NextID()
	if err != nil {
		log.Println(err)
	}

	orderId := strconv.FormatUint(a, 16)
	requestId := strconv.FormatUint(b, 16)

	endpoint := viper.GetString("MOMO_URL")
	partnerCode := viper.GetString("MOMO_PARTNER_CODE")
	accessKey := viper.GetString("MOMO_ACCESS_KEY")
	secretKey := viper.GetString("MOMO_SECRET_KEY")

	amount := strconv.FormatInt(amountV, 10)
	orderInfo := description
	redirectUrl := "http://localhost:8000/api/payment/momoreturn"
	ipnUrl := "http://localhost:8000/api/payment/momonotify"
	requestType := "captureWallet"
	extraData := ""

	// Raw signature string
	rawSignature := fmt.Sprintf("accessKey=%s&amount=%s&extraData=%s&ipnUrl=%s&orderId=%s&orderInfo=%s&partnerCode=%s&redirectUrl=%s&requestId=%s&requestType=%s",
		accessKey, amount, extraData, ipnUrl, orderId, orderInfo, partnerCode, redirectUrl, requestId, requestType)

	signature := utils.GenerateSignature(rawSignature, secretKey)

	payload := MoMoPayload{
		PartnerCode: partnerCode,
		AccessKey:   accessKey,
		RequestID:   requestId,
		Amount:      amount,
		OrderID:     orderId,
		OrderInfo:   orderInfo,
		RedirectUrl: redirectUrl,
		IpnUrl:      ipnUrl,
		ExtraData:   extraData,
		RequestType: requestType,
		Signature:   signature,
		Lang:        "vi",
	}

	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error:", err)
		return "", fiber.NewError(fiber.StatusBadRequest, "Request failed.")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// result := string(body)
	// fmt.Println("MoMo Response:", string(body))

	var requestRes map[string]interface{}
	perr := json.Unmarshal(body, &requestRes)
	if perr != nil {
		fmt.Println("Lỗi parse JSON:", perr)
	}

	// fmt.Printf("Response::::%v", requestRes)

	return requestRes["payUrl"].(string), nil
}
