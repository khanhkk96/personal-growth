package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"personal-growth/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

type MoMoQRPayload struct {
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

func PayViaQRMoMo(amountV int64, description string) (string, error) {
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
	redirectUrl := fmt.Sprintf("%s/api/payment/momo_return", viper.GetString("API_SERVER_ADDRESS"))
	ipnUrl := fmt.Sprintf("%s/api/payment/momo_notify", viper.GetString("API_SERVER_ADDRESS"))
	requestType := "captureWallet"
	extraData := fmt.Sprintf("email=%s", viper.GetString("EMAIL_ADDRESS"))

	// Raw signature string
	rawSignature := fmt.Sprintf("accessKey=%s&amount=%s&extraData=%s&ipnUrl=%s&orderId=%s&orderInfo=%s&partnerCode=%s&redirectUrl=%s&requestId=%s&requestType=%s",
		accessKey, amount, extraData, ipnUrl, orderId, orderInfo, partnerCode, redirectUrl, requestId, requestType)

	signature := utils.CreateMoMoSignature(rawSignature, secretKey)

	payload := MoMoQRPayload{
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

type MomoPayLinkPayload struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	OrderInfo    string `json:"orderInfo"`
	PartnerName  string `json:"partnerName"`
	StoreId      string `json:"storeId"`
	OrderGroupId string `json:"orderGroupId"`
	Lang         string `json:"lang"`
	AutoCapture  bool   `json:"autoCapture"`
	RedirectUrl  string `json:"redirectUrl"`
	IpnUrl       string `json:"ipnUrl"`
	ExtraData    string `json:"extraData"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
}

func PayViaMoMoLink(amountV int64, description string) (string, error) {

	endpoint := viper.GetString("MOMO_URL")
	partnerCode := viper.GetString("MOMO_PARTNER_CODE")
	accessKey := viper.GetString("MOMO_ACCESS_KEY")
	secretKey := viper.GetString("MOMO_SECRET_KEY")

	amount := strconv.FormatInt(amountV, 10)
	orderInfo := description
	redirectUrl := fmt.Sprintf("%s/api/payment/momo_return", viper.GetString("API_SERVER_ADDRESS"))
	ipnUrl := fmt.Sprintf("%s/api/payment/momo_notify", viper.GetString("API_SERVER_ADDRESS"))
	extraData := fmt.Sprintf("email=%s", viper.GetString("EMAIL_ADDRESS"))

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	//randome orderID and requestID
	a, _ := flake.NextID()
	b, _ := flake.NextID()

	var orderId = strconv.FormatUint(a, 16)
	var requestId = strconv.FormatUint(b, 16)
	var partnerName = "MoMo Payment"
	var storeId = "Test Store"
	var orderGroupId = ""
	var autoCapture = true
	var lang = "vi"
	var requestType = "payWithMethod"

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(amount)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)

	signature := utils.CreateMoMoSignature(rawSignature.String(), secretKey)

	var payload = MomoPayLinkPayload{
		PartnerCode:  partnerCode,
		AccessKey:    accessKey,
		RequestID:    requestId,
		Amount:       amount,
		RequestType:  requestType,
		RedirectUrl:  redirectUrl,
		IpnUrl:       ipnUrl,
		OrderID:      orderId,
		StoreId:      storeId,
		PartnerName:  partnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  autoCapture,
		Lang:         lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}

	var jsonPayload []byte
	var err error
	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println("Payload: " + string(jsonPayload))
	// fmt.Println("Signature: " + signature)

	//send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	// fmt.Println("Response from Momo: ", result)

	return result["payUrl"].(string), nil
}

func PayViaVNPay(amountV int64, description string) (string, error) {
	// // config
	vnp_TmnCode := viper.GetString("VNP_TMNCODE")
	vnp_HashSecret := viper.GetString("VNP_HASHSECRET")
	vnp_Url := viper.GetString("VNP_URL")
	vnp_Version := viper.GetString("VNP_VERSION")
	vnp_ReturnUrl := fmt.Sprintf("%s/api/payment/vnpay_return", viper.GetString("API_SERVER_ADDRESS"))

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	b, err := flake.NextID()
	if err != nil {
		log.Println(err)
	}
	orderId := strconv.FormatUint(b, 16)

	now := time.Now()
	txnRef := orderId
	amountStr := fmt.Sprintf("%d", amountV*100) // VNPAY yêu cầu x100

	params := map[string]string{
		"vnp_Version":    vnp_Version,
		"vnp_Command":    "pay",
		"vnp_TmnCode":    vnp_TmnCode,
		"vnp_Amount":     amountStr,
		"vnp_CurrCode":   "VND",
		"vnp_OrderInfo":  description,
		"vnp_OrderType":  "other",
		"vnp_Locale":     "vn",
		"vnp_ReturnUrl":  vnp_ReturnUrl,
		"vnp_IpAddr":     viper.GetString("SERVER_IP"),
		"vnp_CreateDate": now.Format("20060102150405"),
		"vnp_ExpireDate": now.Add(time.Duration(time.Minute * 10)).Format("20060102150405"),
		"vnp_TxnRef":     txnRef,
		// "vnp_TxnRef":     fmt.Sprintf("%s%s", txnRef, now.Format("15:04:05")),
	}

	// Sắp xếp thứ tự param
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Tạo chuỗi dữ liệu để ký
	var hashData strings.Builder
	var query strings.Builder
	for i, k := range keys {
		val := url.QueryEscape(params[k])
		if i > 0 {
			hashData.WriteString("&")
			query.WriteString("&")
		}
		hashData.WriteString(k + "=" + val)
		query.WriteString(k + "=" + val)
	}

	// Ký dữ liệu
	secureHash := utils.CreateVNPayHash(hashData.String(), vnp_HashSecret)

	paymentUrl := fmt.Sprintf("%s?%s&vnp_SecureHash=%s", vnp_Url, query.String(), secureHash)

	return paymentUrl, nil
}
