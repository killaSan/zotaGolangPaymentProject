package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"zotaProblem/definitions"
	"zotaProblem/deposit"
	"zotaProblem/statusCheck"
)

func handleEnvironmentVariables() {
	_, flagMerchantId := os.LookupEnv(definitions.MERCHANT_ID)
	if !flagMerchantId {
		log.Fatal("Merchant id not set as an environment variable")
	}

	_, flagApiSecretKey := os.LookupEnv(definitions.API_SECRET_KEY)
	if !flagApiSecretKey {
		log.Fatal("Api Secret Key not set as an environment variable")
	}

	_, flagCurrency := os.LookupEnv(definitions.CURRENCY)
	if !flagCurrency {
		log.Fatal("Currency not set as an environment variable")
	}

	_, flagEndpointId := os.LookupEnv(definitions.ENDPOINT_ID)
	if !flagEndpointId {
		log.Fatal("Endpoint id not set as an environment variable")
	}

	_, flagBaseUrl := os.LookupEnv(definitions.BASE_URL)
	if !flagBaseUrl {
		log.Fatal("Base URL not set as an environment variable")
	}
}

func main() {
	handleEnvironmentVariables()

	var filePath string
	fmt.Println("Enter filePath")
	fmt.Scan(&filePath)

	var dep deposit.Deposit
	dep.LoadData(filePath)
	dep.SendRequest()

	// If first transaction was not successful there is no point to continue with status chech request
	if dep.Response.Code != "200" {
		fmt.Println("Deposit request was not successful.")
		return
	}
	c := time.Tick(15 * time.Second)
	var statusCheck statusCheck.StatusCheck
	for _ = range c {
		//Download the current contents of the URL and do something with it
		statusCheck.SendRequest(dep)
		if statusCheck.Response.Status == "APPROVED" || statusCheck.Response.Status == "DECLINED" {
			fmt.Println(statusCheck.Response.Status)
			break
		}
		// for card payments we have this error state
		if statusCheck.Response.Status == "ERROR" {
			log.Fatal("Received Error - Order is declined due to a technical error, please inform our support team. final status.")
		}
		if statusCheck.Response.Status == "FILTERED" {
			log.Fatal("Received FILTERED - Order is declined by fraud-prevention system.")
		}
	}
}
