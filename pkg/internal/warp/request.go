package warp

import (
	"fmt"
)

const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
	JSON   ContentType   = "application/json"
)

type RequestMethod string
type ContentType string

type Request struct {
	Warp        string
	Path        string
	ContentType ContentType
	Username    string
	Password    string
}

func (req *Request) Get() ([]byte, error) {
	url := fmt.Sprintf("%s/%s", req.Warp, req.Path)
	return warpCall(url, string(req.ContentType), string(GET), req.Username, req.Password)
}
