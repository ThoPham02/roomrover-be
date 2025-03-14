package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	DataSource string
	RedisCache struct {
		Host     string
		Port     string
		Password string
		DB       int
	}
	ZaloPay struct {
		AppID          string
		Key1           string
		Key2           string
		RedirectDomain string
		BankCode       string
		CallbackUrl    string
	}
}
