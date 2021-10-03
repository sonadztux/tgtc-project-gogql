package model

import (
	"database/sql"

	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/views"
)

func GetCouponByUserID(userID int) (interface{}, error) {
	// variable declarations
	var (
		c   views.Coupon
		err error
		row *sql.Row
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
	WHERE
		coupon_id = $1
	`

	// actual query process
	row = db.QueryRow(query, userID)
	err = row.Scan(&c.CouponID, &c.Name, &c.Category, &c.ImageURL, &c.StartDate, &c.EndDate, &c.Reward)
	if err != nil && err != sql.ErrNoRows {
		return c, err
	}

	return c, err
}
