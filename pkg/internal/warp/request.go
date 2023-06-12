package warp

import (
	"encoding/json"
	"fmt"

	"github.com/HappyTobi/warp/pkg/internal/renderer"
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
	Warp           string
	Path           string
	ContentType    ContentType
	Username       string
	Password       string
	OutputRenderer renderer.Renderer
}

func (req *Request) Get() ([]byte, error) {
	url := fmt.Sprintf("%s/%s", req.Warp, req.Path)
	return warpCall(url, string(req.ContentType), string(GET), req.Username, req.Password, nil)
}

func (req *Request) GetJson() (map[string]interface{}, error) {
	var genJson map[string]interface{}

	data, err := req.Get()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &genJson); err != nil {
		return nil, err
	}

	return genJson, nil
}

func (req *Request) Put(data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", req.Warp, req.Path)
	return warpCall(url, string(req.ContentType), string(PUT), req.Username, req.Password, data)
}

func (req *Request) Post(data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", req.Warp, req.Path)
	return warpCall(url, string(req.ContentType), string(POST), req.Username, req.Password, data)
}
