package conekta

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestAPI returns Conekta API response
func RequestAPI(method string, url string, params ParamsConverter) ([]byte, *Error) {
	dp := defaultParams
	requestURL := BuildURL(url, dp)

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestURL, bytes.NewBuffer(params.Bytes()))

	setHeaders(req, dp)

	res, err := client.Do(req)

	// Client errors
	if err != nil {
		return nil, getConectionError()
	}

	body, _ := ioutil.ReadAll(res.Body)

	// API errors
	if res.StatusCode != 200 {
		return nil, getAPIError(body)
	}

	return body, nil
}

// BuildURL returns base api plus endpoint passed
func BuildURL(url string, p *conektaParams) string {
	return p.apiBase + url
}

// setHeader set req object (htttp.Request) with conekta required headers to auth and identify request
func setHeaders(r *http.Request, p *conektaParams) *http.Request {
	r.Header.Set("Accept", "application/vnd.conekta-v"+p.apiVersion+"+json")
	r.Header.Set("Accept-Language", p.locale)
	r.Header.Set("User-Agent", "Conekta/v1 GoBindings/"+p.version)
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(p.apiKey)))
	r.Header.Set("Content-Type", "application/json")
	return r
}

// New creates a new Conekta POST request,
// it takes a reference struct of conekta's models like
// &conekta.Customer{} and fills it with the response
// if the response has a Conekta's error it returns it and keep empty the Conekta model
func New(v interface{}, p ParamsConverter, endpoint string) *Error {
	res, err := RequestAPI("POST", endpoint, p)
	if err != nil {
		return err
	}
	json.Unmarshal(res, v)
	return nil
}
