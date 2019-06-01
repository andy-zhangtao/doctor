package server

import (
	"github.com/andy-zhangtao/doctor/doctor/model"
	"github.com/graphql-go/graphql"
)

var nodeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "node",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if u, ok := p.Source.(model.RemoteNode); ok {
					return u.Name, nil
				}
				return nil, nil
			},
		},
		"ip": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if u, ok := p.Source.(model.RemoteNode); ok {
					return u.Ip, nil
				}
				return nil, nil
			},
		},
		"comment": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if u, ok := p.Source.(model.RemoteNode); ok {
					return u.Comment, nil
				}
				return nil, nil
			},
		},
	},
})
