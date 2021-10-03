package service

import (
	"errors"

	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/dictionary"
)

func CreateCoupon(data dictionary.Coupon) (*dictionary.Coupon, error) {
	db := database.GetDB()

	query := `
		INSERT INTO coupons (name, coupon_type, status, amount, image_url, usertier, start_date, expire_date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	result, err := db.Exec(query,
		data.Name,
		data.Type,
		data.Status,
		data.Amount,
		data.ImageURL,
		data.UserTier,
		data.StartDate,
		data.ExpireDate,
	)
	
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	return &data, nil
}