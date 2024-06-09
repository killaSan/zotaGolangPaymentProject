package definitions

type ZotaDepositRequest struct {
	MerchantOrderID     string `json:"merchantOrderID"`
	MerchantOrderDesc   string `json:"merchantOrderDesc"`
	OrderAmount         string `json:"orderAmount"`
	OrderCurrency       string `json:"orderCurrency"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerFirstName   string `json:"customerFirstName"`
	CustomerLastName    string `json:"customerLastName"`
	CustomerAddress     string `json:"customerAddress"`
	CustomerCountryCode string `json:"customerCountryCode"`
	CustomerCity        string `json:"customerCity"`
	CustomerZipCode     string `json:"customerZipCode"`
	CustomerPhone       string `json:"customerPhone"`
	CustomerIP          string `json:"customerIP"`
	RedirectUrl         string `json:"redirectUrl"`
	CallbackUrl         string `json:"callbackUrl"`
	CustomParam         string `json:"customParam"`
	CheckoutUrl         string `json:"checkoutUrl"`
	Signature           string `json:"signature"`
}

type ZotaDepositResponseData struct {
	DepositUrl      string `json:"depositUrl"`
	MerchantOrderID string `json:"merchantOrderID"`
	OrderID         string `json:"orderID"`
}

type ZotaDepositResponse struct {
	Code                    string `json:"code"`
	Message                 string `json:"message,omitempty"`
	ZotaDepositResponseData `json:"data,omitempty"`
}

type ZotaOrderStatusCheckRequest struct {
	MerchantID      string `json:"merchantID,omitempty"`
	OrderID         string `json:"orderID,omitempty"`
	MerchantOrderID string `json:"merchantOrderID,omitempty"`
	Timestamp       string `json:"timestamp,omitempty"`
	Signature       string `json:"signature,omitempty"`
}

type ZotaExtraInfoStatusCheckResponse struct {
	AmountChanged     bool   `json:"amountChanged,omitempty"`
	AmountRounded     bool   `json:"amountRounded,omitempty"`
	AmountManipulated bool   `json:"amountManipulated,omitempty"`
	Dcc               bool   `json:"dcc,omitempty"`
	OriginalAmount    string `json:"originalAmount,omitempty"`
	PaymentMethod     string `json:"paymentMethod,omitempty"`
	SelectedBankCode  string `json:"selectedBankCode,omitempty"`
	SelectedBankName  string `json:"selectedBankName,omitempty"`
}

type ZotaDataStatusCheckResponse struct {
	TypeResponse                     string `json:"type"`
	Status                           string `json:"status"`
	ErrorMessage                     string `json:"errorMessage"`
	ProcessorTransactionID           string `json:"processorTransactionID"`
	OrderID                          string `json:"orderID"`
	MerchantOrderID                  string `json:"merchantOrderID"`
	Amount                           string `json:"amount"`
	Currency                         string `json:"currency"`
	CustomerEmail                    string `json:"customerEmail"`
	CustomParam                      string `json:"customParam"`
	ZotaExtraInfoStatusCheckResponse `json:"extraData,omitempty"`
	ZotaOrderStatusCheckRequest      `json:"request,omitempty"`
}

type ZotaOrderStatusCheckResponse struct {
	Code                        string `json:"code"`
	Message                     string `json:"message,omitempty"`
	ZotaDataStatusCheckResponse `json:"data,omitempty"`
}

const (
	MERCHANT_ID    = "MERCHANT_ID"
	API_SECRET_KEY = "API_SECRET_KEY"
	CURRENCY       = "CURRENCY"
	ENDPOINT_ID    = "ENDPOINT_ID"
	BASE_URL       = "BASE_URL"
)
