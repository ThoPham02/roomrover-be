package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf
	CreateBillByTimeJob struct {
		Time      string
		BeforeDay int
	}
}
