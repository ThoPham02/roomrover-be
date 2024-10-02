package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Momo struct {
		Endpoint    string
		AccessKey   string
		SecretKey   string
		OrderInfo   string
		PartnerCode string
		RedirectUrl string
		IpnUrl      string
		PartnerName string
		StoreId     string
		AutoCapture bool
		Lang        string
		RequestType string
	}
}
