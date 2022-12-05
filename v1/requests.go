package dextrade

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	// BaseURL is root API url
	BaseURL = "https://api.dex-trade.com/v1/"
)

// Response is wrapper for standard http.Response and provides
// more methods.
type Response struct {
	Response *http.Response
	Body     []byte
}

// ErrorResponse is the custom error type that is returned if the API returns an
// error.
type ErrorResponse struct {
	Response *Response
	Result   bool   `json:"result"`
	Message  string `json:"error"`
}

// NewRequest create new API request. Relative url can be provided in refURL.
func (c *Client) newRequest(method string, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(BaseURL + "public/" + refURL)
	if err != nil {
		return nil, err
	}
	if params != nil {
		rel.RawQuery = params.Encode()
	}
	u := c.BaseURL.ResolveReference(rel)

	var req *http.Request

	req, err = http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

// newAuthenticatedRequest creates new http request for authenticated routes.
func (c *Client) newAuthenticatedRequest(refURL string, params map[string]string) (*http.Request, error) {
	params["request_id"] = nonce()

	contentForSign := encodeAndSortValues(params)

	sign := signPayload(contentForSign, c.APISecret)

	jsonData, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", BaseURL+"private/"+refURL, bytes.NewBuffer(jsonData))

	req.Header.Add("login-token", c.APIKey)
	req.Header.Add("x-auth-sign", sign)
	req.Header.Add("content-type", "application/json")
	return req, nil
}

func encodeAndSortValues(v map[string]string) string {

	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for key := range v {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		buf.WriteString(vs)
	}
	return buf.String()
}

func (c *Client) performRequest(req *http.Request, v interface{}) (*Response, error) {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = []byte(`Error reading body:` + err.Error())
	}

	response := &Response{resp, body}

	err = checkResponse(response)

	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.Unmarshal(response.Body, v)
		if err != nil {
			return response, err
		}
	}
	return response, nil
}

// checkResponse checks response status code and response
// for errors.
func checkResponse(r *Response) error {
	errorResponse := &ErrorResponse{Response: r}
	err := json.Unmarshal(r.Body, errorResponse)

	if err != nil {
		errorResponse.Message = "Error decoding response error message. " +
			"Please see response body for more information."
	} else if !(errorResponse.Message == "") {
		return errorResponse
	}

	return nil
}

func signPayload(message string, secret string) string {
	message = message + secret

	h := sha256.New()
	h.Write([]byte(message))
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

func nonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Response.Request.Method,
		r.Response.Response.Request.URL,
		r.Response.Response.StatusCode,
		r.Message,
	)
}
