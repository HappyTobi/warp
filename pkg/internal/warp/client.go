package warp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func warpCall(url, contentType, method, username, password string, data []byte) ([]byte, error) {
	req, err := setupRequest(url, contentType, method, data)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		if len(resp.Header["Www-Authenticate"]) > 0 {
			req, err := setupRequest(url, contentType, method, data)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Authorization", digestAuthrization(url, method, username, password, resp))

			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		}
		return nil, fmt.Errorf("could not identify digest authorization")
	case http.StatusOK:
		return io.ReadAll(resp.Body)
	default:
		fmt.Print("Error")
		return nil, fmt.Errorf("Unexpected error")
	}
}

func setupRequest(url, contentType, method string, data []byte) (*http.Request, error) {
	bodyData := bytes.NewBuffer(data)
	hReq, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		return nil, err
	}

	hReq.Header.Set("Content-Type", contentType)

	return hReq, nil
}
