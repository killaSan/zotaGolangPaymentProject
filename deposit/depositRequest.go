package deposit

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"zotaProblem/definitions"
)

type Deposit struct {
	Request  definitions.ZotaDepositRequest
	Response definitions.ZotaDepositResponse
}

const contentType = "application/json"

/**
*	function which will open depositUrl depending on the OS
* 	@param url address of the depositUrl
 */
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (depositRequest *Deposit) SendRequest() {
	postBody := depositRequest.Request.Serialize()
	responseBody := bytes.NewBuffer(postBody)

	address := "https://" + os.Getenv(definitions.BASE_URL) + "/api/v1/deposit/request/"
	address += os.Getenv(definitions.ENDPOINT_ID)

	resp, err := http.Post(address, contentType, responseBody)

	//Handle Error
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	depositRequest.Response.Deserialize(body)
	if depositRequest.Response.Code == "200" {
		openbrowser(depositRequest.Response.DepositUrl)
	}
}

func (depositRequestData *Deposit) LoadData(filePath string) {
	content, err := ioutil.ReadFile(filePath)

	//Handle Error
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &depositRequestData.Request)

	//Handle Error
	if err != nil {
		log.Fatal(err)
	}

	depositRequestData.Request.Signature = depositRequestData.HashSha256(os.Getenv(definitions.ENDPOINT_ID), os.Getenv(definitions.API_SECRET_KEY))
}

func (depositRequestData *Deposit) HashSha256(endpointId string, merchantSecretKey string) string {
	concat := endpointId + depositRequestData.Request.MerchantOrderID + depositRequestData.Request.OrderAmount + depositRequestData.Request.CustomerEmail + merchantSecretKey
	h := sha256.New()
	h.Write([]byte(concat))
	sha256 := h.Sum(nil)
	encode := make([]byte, hex.EncodedLen(len(sha256)))
	hex.Encode(encode, sha256)
	return string(encode)
}
