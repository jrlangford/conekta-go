package conekta

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

// RequestAPI returns Conekta API response
func RequestAPI(method string, url string, params ParamsConverter, customHeaders ...interface{}) ([]byte, error) {
	requestURL := BuildURL(url)

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestURL, bytes.NewBuffer(params.Bytes()))

	setHeaders(req, customHeaders...)

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
func BuildURL(url string) string {
	return apiBase + url
}

// setHeader set req object (htttp.Request) with conekta required headers to auth and identify request
func setHeaders(r *http.Request, customHeaders ...interface{}) *http.Request {
	r.Header.Set("Accept", "application/vnd.conekta-v"+apiVersion+"+json")
	r.Header.Set("Accept-Language", Locale)
	r.Header.Set("User-Agent", "Conekta/v1 GoBindings/"+version)
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(APIKey)))
	r.Header.Set("Content-Type", "application/json")

	for _, v := range customHeaders {
		cht := reflect.TypeOf(v).String()
		if cht == "map[string]string" {
			keys := reflect.ValueOf(v).MapKeys()
			ch := v.(map[string]string)
			for _, k := range keys {
				r.Header.Set(k.String(), ch[k.String()])
			}
		}
	}
	return r
}

// MakeRequest creates a new Conekta request,
// it takes a method, endpoint, parameters and a reference struct of conekta's models like
// &conekta.Customer{} and fills it with the response
// if the response has a Conekta's error it returns it and keep empty the Conekta model
func MakeRequest(method string, endpoint string, p ParamsConverter, v interface{}, customHeaders ...interface{}) error {
	res, err := RequestAPI(method, endpoint, p, customHeaders...)
	if err != nil {
		return err
	}
	json.Unmarshal(res, v)
	return nil
}
