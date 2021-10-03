package views

type Coupon struct {
	ID         int64   `json:"ID"`
	Name       string  `json:"Name"`
	Type       string  `json:"Type"`
	Amount     float64 `json:"Amount"`
	ImageURL   string  `json:"image_url"`
	Tier       string  `json:"usertier"`
	Status     string  `json:"Status"`
	StartDate  string  `json:"start_date"`
	ExpireDate string  `json:"expire_date"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}
