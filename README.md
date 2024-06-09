# zotaGolangPaymentProject
Test Deposit and Status Check for Online Payments

## Before application is run you must set the following environment variables:

1. MERCHANT_ID
2. API_SECRET_KEY
3. USD 
4. ENDPOINT_ID
5. BASE_URL

Example:
export MERCHANT_ID=bigMerchant

export API_SECRET_KEY=apiSecret

export CURRENCY=BGN

export ENDPOINT_ID=123123

export BASE_URL=api.zotapay-stage.com

In case where some of the variables are not set, the program is not run.

## Start application

The Application is started with
go run main.go

1. Enter valid filePath

The User is prompted to enter a valid filePath which will contain our deposit request data. The structure of the file must be a valid JSON. 
Please check deposit/loadDataJson.txt for a valid file content.

2. Deposit Request is send to Zota server. In case of a successful response the deposit url is opened depending on the user OS. 

3. Every 15 seconds a get request is send to Zota server to check the status of the deposit. The polling stops when we get an APPROVED or DECLINED status

4. If we received ERROR or FILTERED status we exit the program. This means some error has occured.