package definitions

import "encoding/json"

type packetUtil interface {
	Serialize() []byte
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

var zeroStatusResponse = &ZotaOrderStatusCheckResponse{}

func (statusCheckResponse *ZotaOrderStatusCheckResponse) Deserialize(jsonStatusCheckResponse []byte) {
	*statusCheckResponse = *zeroStatusResponse
	json.Unmarshal([]byte(jsonStatusCheckResponse), &statusCheckResponse)
}
