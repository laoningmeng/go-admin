package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	viper *viper.Viper
}

func NewConfig() *Conf {
	v := viper.New()
	v.SetConfigName("admin")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &Conf{
		viper: v,
	}
}
