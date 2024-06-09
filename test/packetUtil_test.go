package test

import (
	"testing"
	"zotaProblem/definitions"
	"zotaProblem/deposit"
	"zotaProblem/statusCheck"
)

func TestSerializeJsonDepositRequest(t *testing.T) {
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

	if string(actualJson) != expectedJson {
		t.Error("Struct to JSON doesn't work")
	}

}

// Test to check if sha256 works correctly
func TestHashSha256Deposit(t *testing.T) {
	var dep deposit.Deposit
	dep.Request.MerchantOrderID = "QvE8dZshpKhaOmHY"
	dep.Request.OrderAmount = "50.00"
	dep.Request.CustomerEmail = "customer@email-address.com"
	expectedResult := "cbad0fe994de6cbaff0066077751c222ce310ba4f8d36d7e6adda992ace815df"
	actualResult := dep.HashSha256("433334", "frefre-e3r3r34r-edreferf2d60fb1bf")

	if expectedResult != string(actualResult) {
		t.Error("Hash sha256 doesn't work correctly for Deposit expected ", expectedResult, "but got ", actualResult)
	}
}

func TestHashSha256StatusCheck(t *testing.T) {
	var statusCheck statusCheck.StatusCheck
	statusCheck.Request.MerchantOrderID = "merchOrder123"
	statusCheck.Request.OrderID = "order123"
	statusCheck.Request.Timestamp = "1717961362"
	expectedResult := "6bdedbb6e402c560900e066ee32b96131d72db2b40487f1a0108bb619244868e"
	actualResult := statusCheck.HashSha256("1234", "somesecretkey123")

	if expectedResult != actualResult {
		t.Error("Hash sha256 doesn't work correctly for StatusCheck expected ", expectedResult, "but got ", actualResult)
	}
}
