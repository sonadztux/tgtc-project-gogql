package model

import (
	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/views"
)

func GetAllCoupons() (interface{}, error) {
	// variable declarations
	var (
		coupons []views.Coupon
		err     error
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
			c views.Coupon
		)
		err = rows.Scan(&c.CouponID, &c.Name, &c.Category, &c.ImageURL, &c.StartDate, &c.EndDate, &c.Reward)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, c)
	}

	return coupons, err
}
