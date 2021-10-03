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
		coupons.ID,
		coupons.Name,
		coupons.Type,
		coupons.Status,
		coupons.Amount,
		coupons.image_url,
		coupons.usertier,
		coupons.start_date,
		coupons.expire
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
		err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.ImageURL, &c.StartDate, &c.ExpireDate, &c.Amount, &c.Tier, &c.Status)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, c)
	}

	return coupons, err
}
