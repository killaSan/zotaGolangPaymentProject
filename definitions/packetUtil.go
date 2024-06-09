package definitions

import "encoding/json"

/**
*	Interface which will serializa and deserialize json data from and to a structure
 */
type packetUtil interface {
	/**
	* Serialize method will transform structure to a valid json
	* used for http post requests
	 */
	Serialize() []byte

	/**
	* Deserialize method will take the json response from server and parse it to a structure
	 */
	Deserialize([]byte)
}

func (depositRequest ZotaDepositRequest) Serialize() []byte {
	jsonDepositRequest, _ := json.Marshal(depositRequest)
	return jsonDepositRequest
}

func (depositRequest *ZotaDepositRequest) Deserialize(jsonDepositRequest []byte) {
	json.Unmarshal([]byte(jsonDepositRequest), &depositRequest)
}

func (depositResponse ZotaDepositResponse) Serialize() []byte {
	jsonDepositRequest, _ := json.Marshal(depositResponse)
	return jsonDepositRequest
}

func (depositResponse *ZotaDepositResponse) Deserialize(jsonDepositResponse []byte) {
	json.Unmarshal([]byte(jsonDepositResponse), &depositResponse)
}

func (statusCheckRequest *ZotaOrderStatusCheckRequest) Serialize() []byte {
	jsonStatusCheckRequest, _ := json.Marshal(statusCheckRequest)
	return jsonStatusCheckRequest
}

func (statusCheckRequest *ZotaOrderStatusCheckRequest) Deserialize(jsonStatusCheckRequest []byte) {
	json.Unmarshal([]byte(jsonStatusCheckRequest), &statusCheckRequest)
}

func (statusCheckResponse *ZotaOrderStatusCheckResponse) Serialize() []byte {
	jsonStatusCheckResponse, _ := json.Marshal(statusCheckResponse)
	return jsonStatusCheckResponse
}

// since we poll data every 15 seconds we will need an empty instance of
// statusCheckResponse so that we can clear the previous data
var zeroStatusResponse = &ZotaOrderStatusCheckResponse{}

func (statusCheckResponse *ZotaOrderStatusCheckResponse) Deserialize(jsonStatusCheckResponse []byte) {
	*statusCheckResponse = *zeroStatusResponse
	json.Unmarshal([]byte(jsonStatusCheckResponse), &statusCheckResponse)
}
