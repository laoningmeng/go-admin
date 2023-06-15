package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
	Server ServerConf `mapstructure:"server"`
	Nacos  NacosConf  `mapstructure:"nacos"`
	Consul ConsulConf `mapstructure:"consul"`
}
type ServerConf struct {
	Name         string `mapstructure:"name"`
	Addr         string `mapstructure:"addr"`
	Port         int    `mapstructure:"port"`
	CertFilePath string `mapstructure:"cert_file_path"`
	CertKeyPath  string `mapstructure:"cert_key_path"`
}
type NacosConf struct {
	Addr        string `mapstructure:"addr"`
	Port        int    `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace_id"`
	DataId      string `mapstructure:"data_id"`
	Group       string `mapstructure:"group"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	LogDir      string `mapstructure:"log_dir"`
	CacheDir    string `mapstructure:"cache_dir"`
	Timeout     int    `mapstructure:"timeout"`
	LogLevel    string `mapstructure:"log_level"`
}
type ConsulConf struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Token string `mapstructure:"token"`
}

type MysqlConn struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

// 业务conf
type ServiceConf struct {
	Mysql MysqlConn
}

func (c *Conf) GetConf() *Conf {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("yml")
	path, _ := os.Getwd()
	fmt.Println("path:", path)
	v.AddConfigPath("../")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var res Conf
	err = v.Unmarshal(&res)
	if err != nil {
		panic(err)
	}
	return &res
}

func GetConsulConf() ConsulConf {
	conf := Conf{}
	return conf.GetConf().Consul
}
func GetServer() ServerConf {
	conf := Conf{}
	return conf.GetConf().Server
}
func GetNacosConf() NacosConf {
	conf := Conf{}
	return conf.GetConf().Nacos
}
