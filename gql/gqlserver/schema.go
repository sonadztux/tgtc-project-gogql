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
							Type: graphql.EnumValueType,
						},
						"status": &graphql.ArgumentConfig{
							Type: graphql.EnumValueType,
						},
						"amount": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"image_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"usertier": &graphql.ArgumentConfig{
							Type: graphql.EnumValueType,
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.DateTime,
						},
						"expire_date": &graphql.ArgumentConfig{
							Type: graphql.DateTime,
						},
					},
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