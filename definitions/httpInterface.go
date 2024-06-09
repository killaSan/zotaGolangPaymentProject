package definitions

// Interface used for Post Requests
type HttpPost interface {
	/**
	* method which creates secret key to pass along post requests
	* @return: string
	**/
	HashSha256(endpointId string, merchantSecretKey string) string
	/**
	* method which loads file content to deposit request structure.
	* @param - filename which should be a valid filepath. It's structure should be a valid JSON
	**/
	LoadData(string)

	/**
	* method that sends http post request and the return value is loaded to response
	**/
	SendRequest()
}

type HttpGet interface {
	/**
	* method which creates secret key to pass along get requests
	* @return: string
	**/
	HashSha256(merchantId string, merchantSecretKey string) string
	/**
	* method which creates http get request sends it to zota server
	* The returned http response is loaded to reponse structure
	**/
	SendRequest(HttpPost)
}
