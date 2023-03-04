package warp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func warpCall(url, contentType, method, username, password string) ([]byte, error) {
	hReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	hReq.Header.Set("Content-Type", contentType)

	resp, err := client.Do(hReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		if len(resp.Header["Www-Authenticate"]) > 0 {
			hReq.Header.Set("Authorization", digestAuthrization(url, method, username, password, resp))
			resp, err := client.Do(hReq)
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)
		}
		return nil, fmt.Errorf("could not identify digest authorization")
	case http.StatusOK:
		return ioutil.ReadAll(resp.Body)
	default:
		fmt.Print("Error")
		return nil, fmt.Errorf("Unexpected error")
	}
}
