package gqlserver

import "github.com/graphql-go/graphql"

var CouponType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:	"Coupon",
		Description: "Detail of the coupon",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"usertier": &graphql.Field{
				Type: graphql.String,
			},
			"start_date": &graphql.Field{
				Type: graphql.String,
			},
			"expire_date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)