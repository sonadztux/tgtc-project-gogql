package gqlserver

import "github.com/graphql-go/graphql"

type SchemaWrapper struct {
	couponResolver	*Resolver
	Schema 			graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithCouponResolver(coupon *Resolver) *SchemaWrapper {
	s.couponResolver = coupon
	return s
}

func (s *SchemaWrapper) Init() error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "CouponsGetter",
			Description: "All query related to getting coupon data",
			Fields: graphql.Fields{
				"CouponUser": &graphql.Field{
					Type:        CouponType,
					Description: "Get coupon by user ID",
					Args: graphql.FieldConfigArgument{
						"UserID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.GetCouponByUserID(),
				},
				"Coupons": &graphql.Field{
					Type:        graphql.NewList(CouponType),
					Description: "Get all coupons",
					Resolve:     s.couponResolver.GetAllCoupons(),
				},
			},
		}),
		
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:		"CouponSetter",
			Description: "All query related to create or modify coupon data",
			Fields: graphql.Fields{
				"CreateCoupon": &graphql.Field{
					Type: graphql.Boolean,
					Description: "Create coupon",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"type": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"status": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"amount": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"image_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"usertier": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"expire_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.couponResolver.CreateCoupon(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}