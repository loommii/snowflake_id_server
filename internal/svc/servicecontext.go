package svc

import (
	"snowflake_id_server/internal/config"

	"github.com/loommii/flakegen"
)

type ServiceContext struct {
	Config       config.Config
	FlakegenNode *flakegen.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	flakegenNode, err := flakegen.NewNode(1, 1)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		FlakegenNode: flakegenNode,
	}
}
