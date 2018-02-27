package conekta

//
var APIKey string

//
var Locale string

type conektaParams struct {
	apiKey     string
	apiBase    string
	apiVersion string
	locale     string
	version    string
}

var defaultParams = &conektaParams{
	apiKey:     APIKey,
	locale:     Locale,
	apiBase:    "https://api.conekta.io",
	apiVersion: "2.0.0",
	version:    "0.0.1",
}
