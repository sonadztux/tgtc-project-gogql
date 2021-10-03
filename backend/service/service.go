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

func GetAllCoupons() ([]dictionary.Coupon, error) {
	// get current database connection
	db := database.GetDB()

	// construct sql statement
	query := `
	SELECT
		id,
		name,
		coupon_type,
		status,
		amount,
		image_url,
		usertier,
		start_date,
		expire_date
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
	var result []dictionary.Coupon
	for rows.Next() {
		var data dictionary.Coupon
		rows.Scan(
			&data.ID,
			&data.Name,
			&data.Type, 
			&data.Status,
			&data.Amount,
			&data.ImageURL,
			&data.UserTier,
			&data.StartDate,
			&data.ExpireDate,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, err
}

func GetCouponByUserID(userID int) (*dictionary.Coupon, error) {
	// get current database connection
	db := database.GetDB()

	// construct sql statement
	query := `
	SELECT
		id,
		name,
		coupon_type,
		status,
		amount,
		image_url,
		usertier,
		start_date,
		expire_date
	FROM
		coupons
	JOIN
		detail_coupon
	ON
		detail_coupon.coupon_id = coupons.id
	WHERE
		detail_coupon.user_id = $1
	`

	// actual query process
	row := db.QueryRow(query, userID)

	var result dictionary.Coupon
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Type, 
		&result.Status,
		&result.Amount,
		&result.ImageURL,
		&result.UserTier,
		&result.StartDate,
		&result.ExpireDate,
	)
	if err != nil {
		return nil, err
	}

	return &result, err
}