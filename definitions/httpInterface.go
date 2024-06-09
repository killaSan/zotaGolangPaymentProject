package definitions

// Interface used for Post Requests
type HttpPost interface {
	/**
	* method which creates secret key  to pass along post requests
	* @param: endpointId which is set as an environment variable before application is started
	* @param: merchantSecretKey which is set as an environment variable before application is started
	* @return: encodedstring
	 */
	HashSha256(endpointId string, merchantSecretKey string) string
	/**
	* method which loads file content to deposit request structure. File content must be a valid JSON
	* @param - filename which should be a valid filepath.
	 */
	LoadData(string)

	/**
	* method that sends http post request
	 */
	SendRequest()
}

// Interface used for get requests
type HttpGet interface {
	/**
	* method which creates secret key to pass along get requests
	* @param merchantId - this is an environment variable which should be set before the application is run
	* @param merchantSecretKey - this is an en environment variable which should be set before the application is run
	* @return: string
	 */
	HashSha256(merchantId string, merchantSecretKey string) string
	/**
	* method which creates http get request sends it to zota server
	* @param: httpPost interface from which we read the response data and load it to our interface request
	 */
	SendRequest(HttpPost)
}
