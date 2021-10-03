package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/sonadztux/tgtc-project-gogql/backend/dictionary"
	"github.com/sonadztux/tgtc-project-gogql/backend/service"
)

type Resolver struct {}

func NewResolver() *Resolver {
	return &Resolver{}
}

// func (r *Resolver) AllCoupons() graphql.FieldResolveFn {
// 	return func(c graphql.ResolveParams) (interface{}, error) {
// 		id, _ := c.Args["id"].(int)
// 		return service.GetAllCoupons(id)
// 	}
// }

func (r *Resolver) CreateCoupon() graphql.FieldResolveFn {
	return func(c graphql.ResolveParams) (interface{}, error) {
		name, _ := c.Args["name"].(string)
		coupontype, _ := c.Args["type"].(string)
		status, _ := c.Args["status"].(string)
		amount, _ := c.Args["amount"].(float64)
		image_url, _ := c.Args["image_url"].(string)
		usertier, _ := c.Args["usertier"].(string)
		start_date, _ := c.Args["start_date"].(string)
		expire_date, _ := c.Args["expire_date"].(string)

		req := dictionary.Coupon{
			Name:		name,
			Type:		coupontype,
			Status:		status,
			Amount:		amount,
			ImageURL:	image_url,
			UserTier:	usertier,
			StartDate:	start_date,
			ExpireDate:	expire_date,
		}

		_, err := service.CreateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
