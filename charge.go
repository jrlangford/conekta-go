package conekta

//ChargeParams is the object to fill after api call
type ChargeParams struct {
	PaymentMethod *PaymentMethodParams `json:"payment_method,omitempty"`
}

//PaymentMethodParams is the object to fill after api call
type PaymentMethodParams struct {
	Type                string `json:"type,omitempty"`
	TokenID             string `json:"token_id,omitempty"`
	PaymentSourceID     string `json:"payment_source_id,omitempty"`
	MonthlyInstallments int64  `json:"monthly_installments,omitempty"`
	ExpiresAt           int64  `json:"expires_at,omitempty"`
}

//Charge should be a struct of the api response
type Charge struct {
	ID                  string          `json:"id,omitempty"`
	Object              string          `json:"object,omitempty"`
	Status              string          `json:"status,omitempty"`
	Amount              int64           `json:"amount,omitempty"`
	OrderID             string          `json:"order_id,omitempty"`
	CustomerID          string          `json:"customer_id,omitempty"`
	Livemode            bool            `json:"livemode,omitempty"`
	MonthlyInstallments int64           `json:"monthly_installments,omitempty"`
	CreatedAt           int64           `json:"created_at,omitempty"`
	PaidAt              int64           `json:"paid_at,omitempty"`
	Currency            string          `json:"currency,omitempty"`
	Description         string          `json:"description,omitempty"`
	PaymentMethod       *PaymentMethod  `json:"payment_method,omitempty"`
}

//PaymentMethod should be a struct of the api response
type PaymentMethod struct {
	ID                  string          `json:"id,omitempty"`
	Object              string          `json:"object,omitempty"`
	Type                string          `json:"type,omitempty"`
	CreatedAt           int64           `json:"created_at,omitempty"`
	Last4               string          `json:"last4,omitempty"`
	Bin                 string          `json:"bin,omitempty"`
	ExpMonth            string          `json:"exp_month,omitempty"`
	ExpYear             string          `json:"exp_year,omitempty"`
	Brand               string          `json:"brand,omitempty"`
	Name                string          `json:"name,omitempty"`
	ParentID            string          `json:"parent_id,omitempty"`
	Default             bool            `json:"default,omitempty"`
	Reference           string          `json:"reference,omitempty"`
	StoreName           string          `json:"store_name,omitempty"`
	ServiceName         string          `json:"service_name,omitempty"`
	ServiceNumber       string          `json:"service_number,omitempty"`
	ExpiresAt           int64           `json:"expires_at,omitempty"`
	Description         string          `json:"description,omitempty"`
	AuthCode            int64           `json:"auth_code,omitempty"`
	Store               string          `json:"store,omitempty"`
	Address             *DefaultAddress `json:"address,omitempty"`
}

//DefaultAddress should be nested struct of PaymentMethod
type DefaultAddress struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	Object     string `json:"object,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

//ChargesList is a list of shippingLines
type ChargesList struct {
	ListMeta
	Data []*Charge `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *ChargeParams) Bytes() []byte {
	return paramsToBytes(c)
}
