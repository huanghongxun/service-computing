package schema

type HTTPError struct {
	Error HTTPErrorItem `json:"error"`
}

type HTTPErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPStatus struct {
	Status string `json:"status"`
}
