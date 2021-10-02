package model

import (
	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/views"
)

func GetAllCoupons() (interface{}, error) {
	// variable declarations
	var (
		products []views.Coupon
		err      error
	)
	// get current database connection
	db := database.GetDB()

	// construct sql statement
	query := `
	SELECT
		coupon_id,
		coupon_name,
		category,
		image_url,
		start_date,
		end_date,
		reward
	FROM
		coupons
	`

	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// loop and struct scan
	for rows.Next() {
		var (
			p views.Coupon
		)
		err = rows.Scan(&p.CouponID, &p.Name, &p.Category, &p.ImageURL, &p.StartDate, &p.EndDate, &p.Reward)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, err
}
