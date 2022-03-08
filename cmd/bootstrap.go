package cmd

import (
	pool "git.yongche.com/rabbitmq-channel/pkg"
)

type bootStrap struct {
	PoolService     *pool.PoolService
}

func newBootStrap(options *ApplicationOptions) *bootStrap {
	return &bootStrap{
		PoolService: pool.NewPoolService(options.PoolOption),
	}
}
