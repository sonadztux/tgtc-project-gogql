package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/sonadztux/tgtc-project-gogql/backend/model"
)

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}

// func (r *Resolver) AllCoupons() graphql.FieldResolveFn {
// 	return func(c graphql.ResolveParams) (interface{}, error) {
// 		id, _ := c.Args["id"].(int)
// 		return service.GetAllCoupons(id)
// 	}
// }

// func (r *Resolver) CreateCoupon() graphql.FieldResolveFn {

// }

func (r *Resolver) GetCouponByUserID() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)

		// update to use Usecase from previous session
		return model.GetCouponByUserID(id)
	}
}

func (r *Resolver) GetAllCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return model.GetAllCoupons()
	}
}
