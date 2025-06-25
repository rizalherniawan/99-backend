package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rizalherniawan/99-backend-test/public-api/internal/common"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
)

func CallOtherApi(httpMethod string, path string, key string, payload io.Reader, form url.Values) (*http.Response, error) {
	client := &http.Client{}
	var req *http.Request

	// applies when the API requires form
	if form != nil {
		tempReq, _ := http.NewRequest(httpMethod, path, strings.NewReader(form.Encode()))
		req = tempReq
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {

		// applies form is not required
		tempReq, _ := http.NewRequest(httpMethod, path, payload)
		req = tempReq
	}

	// set API key to header
	req.Header.Set("X-API-KEY", key)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("failed to call due to: ", err.Error())
		return nil, err
	}
	return resp, nil
}

func ExtractResponse(resp *http.Response) ([]byte, *exception.ApiError) {

	// applies when the API throws error code
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var errResponse common.ErrResponseDto
		if err := json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			log.Println("failed to extract json due to: ", err.Error())
			return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
		}
		return nil, exception.NewApiError(errResponse.Errors, resp.StatusCode)
	}
	defer resp.Body.Close()

	// extract the response body
	body, er := io.ReadAll(resp.Body)
	if er != nil {
		log.Println("failed to extract json due to: ", er.Error())
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	return body, nil
}
