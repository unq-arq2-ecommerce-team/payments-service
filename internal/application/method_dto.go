package application

type PaymentMethodDto struct {
	Type    string                 `json:"type"`
	Details map[string]interface{} `json:"details"`
}
