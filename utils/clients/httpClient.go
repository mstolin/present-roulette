package clients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPFacade struct {
	client http.Client
}

// Sends a request to the given url.
// If the request is a GET or DELETE request, data can be nil.
func (facade HTTPFacade) do(method string, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := facade.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	// A non 2xx status does not cause an error
	if resp.StatusCode != 200 {
		// TODO Hier kann man ganz einfach die fehlermeldung weiterleiten!!
		/*errResp := httpErrors.ErrorResponse{}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return body, err
		}*/
		return []byte{}, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return body, nil
}

// Constructs a new HTTPFacade instance.
func NewHTTPFacade() HTTPFacade {
	facade := HTTPFacade{}
	facade.client = http.Client{}
	return facade
}

// Sends a POST request to the given url.
func (facade HTTPFacade) DoPost(url string, data []byte) ([]byte, error) {
	return facade.do("POST", url, data)
}

// Sends a GET request to the given url.
func (facade HTTPFacade) DoGet(url string) ([]byte, error) {
	return facade.do("GET", url, nil)
}

// Sends a PUT request to the given url.
func (facade HTTPFacade) DoPut(url string, data []byte) ([]byte, error) {
	return facade.do("PUT", url, data)
}

// Sends a DELETE request to the given url.
func (facade HTTPFacade) DoDelete(url string) ([]byte, error) {
	return facade.do("DELETE", url, nil)
}
