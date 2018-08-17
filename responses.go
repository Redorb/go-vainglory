package vainglory

// StatusResponse is the response payload for the status end point
type StatusResponse struct {
	Data StatusResponseData `json:"data"`
}

// StatusResponseData contains all of the data returned in the StatusResponse
type StatusResponseData struct {
	typeIDPair
	Attributes StatusAttributes `json:"attributes"`
}

// StatusAttributes contains all of the attributes returned in the StatusResponse
type StatusAttributes struct {
	Released string `json:"releasedAt"`
	Version  string `json:"version"`
}

// typeIDPair is a common pattern used throughout all responses
type typeIDPair struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
