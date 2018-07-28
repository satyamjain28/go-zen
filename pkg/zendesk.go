package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// NewRequest calls the zendesk API for the provided function
func NewRequest(auth Access, function string, method string, payload []byte, expectedStatusCode int) ([]byte, ErrorResponse) {
	requestURL := fmt.Sprintf("https://%s.zendesk.com/api/v2/%s", auth.Domain, function)

	req, err := http.NewRequest(method, requestURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error in creating the request")
		return nil, ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(auth.UserName, auth.Password)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Printf("Error in the api call for the function %s", function)
		return nil, ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}

	if res.StatusCode == 401 {
		log.Printf("Authenitcation failed")
		return nil, ErrorResponse{Message: "error in authenticating the user with zendesk", Code: 401, IsError: true}
	}

	if res.StatusCode != expectedStatusCode {
		log.Printf("API call failed with status code as %d", res.StatusCode)
		return nil, ErrorResponse{Message: "api call failed", Code: res.StatusCode, IsError: true}
	}

	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("Error in unmarshalling the body of the response \n%v", err)
		return nil, ErrorResponse{Message: err.Error(), Code: res.StatusCode, IsError: true}
	}
	defer res.Body.Close()

	return contents, ErrorResponse{IsError: false}
}
