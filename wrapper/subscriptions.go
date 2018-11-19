package wrapper

/********** GLOBAL VARIABLES **********/
const subsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *ClientCredentials) GetSubscriptions() map[string]interface{} {
	res, body, errs := request.
		Get(subsURL).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "subscriptions")
}

// GetSubscription returns a single subscription
func (c *ClientCredentials) GetSubscription(subID string) map[string]interface{} {
	url := subsURL + "/" + subID

	res, body, errs := request.
		Get(url).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "subscriptions")
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *ClientCredentials) CreateSubscription(data string) map[string]interface{} {
	res, body, errs := request.
		Post(subsURL).
		Set("x-sp-gateway", c.gateway).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "subscriptions")
}

// UpdateSubscription updates an existing subscription
func (c *ClientCredentials) UpdateSubscription(subID string, data string) map[string]interface{} {
	header(c, gatewaySetting)

	url := subsURL + "/" + subID

	res, body, errs := request.
		Put(url).
		Set("x-sp-gateway", c.gateway).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "subscriptions")
}
