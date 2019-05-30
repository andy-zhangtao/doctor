package server

import (
	"github.com/andy-zhangtao/doctor/doctor/model"
	"github.com/andy-zhangtao/doctor/doctor/node"
	"github.com/andy-zhangtao/doctor/doctor/store"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

var addUser = &graphql.Field{
	Type:        graphql.String,
	Description: "Add New Doctor User",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)

		logrus.Debugf("Create New User [%s]\n", name)

		err := node.InitDoctorNode(name)
		if err != nil {
			return nil, err
		}

		return "Succ", nil
	},
}

var deleteNode = &graphql.Field{
	Type:        graphql.String,
	Description: "Delete Node Info",
	Args: graphql.FieldConfigArgument{
		"ip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		ip, _ := p.Args["ip"].(string)

		rn := model.RemoteNode{
			Ip: ip,
		}

		err := node.DeleteNode(rn)
		if err != nil {
			return nil, err
		}

		return "Delete Succ", nil
	},
}

var updateNode = &graphql.Field{
	Type:        graphql.String,
	Description: "Update Node Info",
	Args: graphql.FieldConfigArgument{
		"ip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"comment": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		comment, _ := p.Args["comment"].(string)
		ip, _ := p.Args["ip"].(string)

		rn := model.RemoteNode{
			Ip:      ip,
			Comment: comment,
		}

		err := node.UpdateNodeComment(rn)
		if err != nil {
			return nil, err
		}

		return "Update Succ", nil
	},
}
var watchNodes = &graphql.Field{
	Type:        graphql.NewList(nodeType),
	Description: "Get Node Info",
	Args: graphql.FieldConfigArgument{
		"ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ip, _ := p.Args["ip"].(string)
		var rns []model.RemoteNode

		if ip != "" {
			_rn := model.RemoteNode{
				Ip: ip,
			}
			err := node.FindOneNode(&_rn)
			if err != nil {
				return nil, err
			}

			rns = append(rns, _rn)
		} else {
			var err error
			rns, err = node.FindAllNodes()
			if err != nil {
				return nil, err
			}
		}

		return rns, nil
	},
}

var addNode = &graphql.Field{
	Type:        graphql.String,
	Description: "Add A New Node",
	Args: graphql.FieldConfigArgument{
		"comment": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"ip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		comment, _ := params.Args["comment"].(string)
		ip, _ := params.Args["ip"].(string)
		user, _ := params.Args["user"].(string)
		password, _ := params.Args["password"].(string)

		rn := model.RemoteNode{
			Ip:       ip,
			Name:     user,
			Password: password,
			Comment:  comment,
		}

		err := node.NodeInit(rn)
		if err != nil {
			return nil, err
		}

		err = store.SaveRemoteNode(rn)
		if err != nil {
			return nil, err
		}

		return "Node Init Succ", nil
	},
}
