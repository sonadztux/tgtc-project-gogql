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
		coupons.ID,
		coupons.Name,
		coupons.Type,
		coupons.Status,
		coupons.Amount,
		coupons.image_url,
		coupons.usertier,
		coupons.start_date,
		coupons.expire,
	FROM
		coupons
	JOIN
		DetailCoupon
	ON
		DetailCoupon.CouponID = coupons.ID
	WHERE
		DetailCoupon.UserID = $1
	`

	// actual query process
	row = db.QueryRow(query, userID)
	err = row.Scan(&c.ID, &c.Name, &c.Type, &c.ImageURL, &c.StartDate, &c.ExpireDate, &c.Amount, &c.Tier, &c.Status)
	if err != nil && err != sql.ErrNoRows {
		return c, err
	}

	return c, err
}
