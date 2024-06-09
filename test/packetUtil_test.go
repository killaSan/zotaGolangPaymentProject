package test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"zotaProblem/definitions"
	"zotaProblem/deposit"
)

func TestSerializeJson(t *testing.T) {
	depositRequest := definitions.ZotaDepositRequest{
		MerchantOrderID:     "QvE8dZshpKhaOmHY",
		MerchantOrderDesc:   "Test order",
		OrderAmount:         "500.00",
		OrderCurrency:       "THB",
		CustomerEmail:       "customer@email-address.com",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan",
		CustomerCountryCode: "TH",
		CustomerCity:        "Surat Thani",
		CustomerZipCode:     "84280",
		CustomerPhone:       "+66-77999110",
		CustomerIP:          "103.106.8.104",
		RedirectUrl:         "https://www.example-merchant.com/payment-return/",
		CallbackUrl:         "https://www.example-merchant.com/payment-callback/",
		CustomParam:         "{\"UserId\": \"e139b447\"}",
		CheckoutUrl:         "https://www.example-merchant.com/account/deposit/?uid=e139b447",
		Signature:           "47d7ed292cf10e689b311ef5573eddbcc8505fe51e20d3f74e6b33756d96800b",
	}

	var expectedJson string = `{"merchantOrderID":"QvE8dZshpKhaOmHY","merchantOrderDesc":"Test order","orderAmount":"500.00","orderCurrency":"THB","customerEmail":"customer@email-address.com","customerFirstName":"John","customerLastName":"Doe","customerAddress":"5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan","customerCountryCode":"TH","customerCity":"Surat Thani","customerZipCode":"84280","customerPhone":"+66-77999110","customerIP":"103.106.8.104","redirectUrl":"https://www.example-merchant.com/payment-return/","callbackUrl":"https://www.example-merchant.com/payment-callback/","customParam":"{\"UserId\": \"e139b447\"}","checkoutUrl":"https://www.example-merchant.com/account/deposit/?uid=e139b447","signature":"47d7ed292cf10e689b311ef5573eddbcc8505fe51e20d3f74e6b33756d96800b"}`

	actualJson := depositRequest.Serialize()
	fmt.Println(actualJson)
	fmt.Println(expectedJson)
	if string(actualJson) != expectedJson {
		t.Error("Struct to JSON doesn't work")
	}
}

// Test to check if sha256 works correctly
func TestHashSha256(t *testing.T) {
	var dep deposit.Deposit
	dep.Request.MerchantOrderID = "QvE8dZshpKhaOmHY"
	dep.Request.OrderAmount = "50.00"
	dep.Request.CustomerEmail = "customer@email-address.com"
	expectedResult := "e4d9fd4157d33c803597ca6a830c3a46b020792b6ff6c0bc8160618cc613ed5e"
	hassha256 := dep.HashSha256("402334", "5f4a6fcf-9048-4a0b-afc2-ed92d60fb1bf")
	encode := make([]byte, hex.EncodedLen(len(hassha256)))
	hex.Encode(encode, hassha256)
	actualResult := string(encode)
	if expectedResult != string(actualResult) {
		t.Error("Hash sha256 doesn't work correctly expected ", expectedResult, "but got ", actualResult)
	}
}
