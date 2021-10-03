package dictionary

type Coupon struct {
	ID			int64	`json:"id"`
	Name		string	`json:"name"`
	Type 		string	`json:"type"`
	Status 		string	`json:"status"`
	Amount		float64	`json:"amount"`
	ImageURL	string	`json:"image_url"`
	UserTier	string	`json:"usertier"`
	StartDate	string	`json:"start_date"`
	ExpireDate	string	`json:"expire_date"`
}