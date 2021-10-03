package views

type Coupon struct {
	CouponID  int64   `json:"coupon_id"`
	Name      string  `json:"coupon_name"`
	Category  string  `json:"category"`
	Reward    float64 `json:"reward"`
	ImageURL  string  `json:"image_url"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}
