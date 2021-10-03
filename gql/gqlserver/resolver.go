package gqlserver

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

// func (r *Resolver) CreateCoupon() graphql.FieldResolveFn {

// }