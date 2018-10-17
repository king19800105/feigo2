package io

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SMSSendRequest collects the request parameters for the SMSSend method.
type SMSSendRequest struct {
	Username      string `json:"username"`
	Key           string `json:"key"`
	Sign          string `json:"sign"`
	Content       string `json:"content"`
	ExtensionCode string `json:"extensionCode"`
}

// SMSQueryRequest collects the request parameters for the SMSQuery method.
type SMSQueryRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

// SMSSendResponse collects the response parameters for the SMSSend method.
type ApiResponse struct {
	Err    error  `json:"message"`
	Result string `json:"result"`
}

// Failed implements Failer.
func (r ApiResponse) Failed() error {
	return r.Err
}
