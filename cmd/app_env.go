package cmd

import (
	pool "git.yongche.com/rabbitmq-channel/pkg"
	"github.com/spf13/viper"
	"log"
)

type ApplicationOptions struct {
	PoolOption *pool.Options `mapstructure:"pool_option"`
}

func (a *ApplicationOptions) Load() {
	err := viper.Unmarshal(a)
	if err != nil {
		log.Fatalln(err)
	}
}

