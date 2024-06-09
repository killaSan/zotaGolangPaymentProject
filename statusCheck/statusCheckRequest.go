package statusCheck

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"zotaProblem/definitions"
	"zotaProblem/deposit"
)

// status Check structure - holds the http get request and response
type StatusCheck struct {
	Request  definitions.ZotaOrderStatusCheckRequest
	Response definitions.ZotaOrderStatusCheckResponse
}

const contentType = "application/json"

func (statusCheckRequestData *StatusCheck) SendRequest(dep deposit.Deposit) {
	statusCheckRequestData.Request.MerchantID = os.Getenv(definitions.MERCHANT_ID)
	statusCheckRequestData.Request.MerchantOrderID = dep.Request.MerchantOrderID
	statusCheckRequestData.Request.OrderID = dep.Response.OrderID
	statusCheckRequestData.Request.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	statusCheckRequestData.Request.Signature = statusCheckRequestData.HashSha256(os.Getenv(definitions.MERCHANT_ID), os.Getenv(definitions.API_SECRET_KEY))

	qs := "?merchantID=" + statusCheckRequestData.Request.MerchantID + "&merchantOrderID=" + statusCheckRequestData.Request.MerchantOrderID + "&orderID=" + statusCheckRequestData.Request.OrderID + "&timestamp=" + statusCheckRequestData.Request.Timestamp + "&signature=" + statusCheckRequestData.Request.Signature
	var address string = "https://" + os.Getenv(definitions.BASE_URL) + "/api/v1/query/order-status/"
	requestURL := address + qs
	resp, err := http.Get(requestURL)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	statusCheckRequestData.Response.Deserialize(body)
}

func (statusCheck StatusCheck) HashSha256(merchantId string, merchantSecretKey string) string {
	concat := merchantId + statusCheck.Request.MerchantOrderID + statusCheck.Request.OrderID + statusCheck.Request.Timestamp + merchantSecretKey

	h := sha256.New()
	h.Write([]byte(concat))
	sha256 := h.Sum(nil)
	encode := make([]byte, hex.EncodedLen(len(sha256)))
	hex.Encode(encode, sha256)
	return string(encode)
}
