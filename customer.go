package conekta

import "encoding/json"

// Customer represents a Conekta customer.
// For details see https://developers.conekta.com/api#customer.
type Customer struct {
	ID                       string `json:"id"`
	Object                   string `json:"object"`
	Live                     bool   `json:"livemode"`
	CreatedAt                int64  `json:"created_at"`
	Name                     string `json:"name"`
	Email                    string `json:"email"`
	Phone                    string `json:"phone"`
	Corporate                bool   `json:"corporate"`
	DefaultShippingContactID string `json:"default_shipping_contact_id"`
	DefaultPaymentSourceID   string `json:"default_payment_source_id"`
	PaymentSources           string `json:"payment_sources"`
	ShippingContacts         string `json:"shipping_contacts"`
	Subscription             string `json:"subscription"`
}

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For details see https://developers.conekta.com/api#create-customer and https://developers.conekta.com/api#update-customer.
type CustomerParams struct {
	Name             string                   `json:"name"`
	Phone            string                   `json:"phone"`
	Email            string                   `json:"email"`
	PaymentSources   []PaymentSourcesParams   `json:"payment_sources"`
	Corporate        bool                     `json:"corporate"`
	ShippingContacts []ShippingContactsParams `json:"shipping_contacts"`
}

// PaymentSourcesParams TODO separate this submodel
type PaymentSourcesParams struct {
	TokenID     string `json:"token_id"`
	PaymentType string `json:"type"`
}

// ShippingContactsParams TODO separate this submodel
type ShippingContactsParams struct {
	Phone    string    `json:"phone"`
	Receiver string    `json:"receiver"`
	Address  []Address `json:"address"`
}

// Address TODO separate this submodel
type Address struct {
	Street1    string `json:"street1"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

// CustomerList is a list of customers
type CustomerList []*Customer

// ParamsConverter is an interface to handle Conekta params conversions
type ParamsConverter interface {
	Bytes() []byte
}

func paramsToBytes(v interface{}) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		return []byte{}
	}
	return r
}

// Bytes converts a CustomerParams to []byte
func (c *CustomerParams) Bytes() []byte {
	return paramsToBytes(c)
}
