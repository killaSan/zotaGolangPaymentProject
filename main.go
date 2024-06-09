package main

import (
	"fmt"
	"os"
	"time"
	"zotaProblem/deposit"
	"zotaProblem/statusCheck"
)

var mecrchantID string
var aPISecretKey string
var currency string
var endpointID string
var baseUrl string

func setEnvVariables(readValueFromUser string, promptMessage string, environmentName string) {
	fmt.Println(promptMessage)
	fmt.Scan(&readValueFromUser)
	os.Setenv(environmentName, readValueFromUser)
}

func readInput() {
	//setEnvVariables(mecrchantID, "Enter MerchantId", definitions.MERCHANT_ID)
	//setEnvVariables(aPISecretKey, "Enter API Secret Key", definitions.API_SECRET_KEY)
	//	setEnvVariables(currency, "Enter Currency", definitions.CURRENCY)
	//	setEnvVariables(endpointID, "Enter EndPointID", definitions.ENDPOINT_ID)
	//	setEnvVariables(baseUrl, "Enter Base URL", definitions.BASE_URL)
}

func main() {
	var dep deposit.Deposit

	readInput()
	var filePath string
	fmt.Println("Enter filePath")
	fmt.Scan(&filePath)
	dep.LoadData(filePath)
	dep.SendRequest()
	c := time.Tick(15 * time.Second)

	var statusCheck statusCheck.StatusCheck

	for _ = range c {
		//Download the current contents of the URL and do something with it
		statusCheck.SendRequest(dep)
		if statusCheck.Response.Status == "APPROVED" || statusCheck.Response.Status == "DECLINED" {
			break
		}
	}
}
