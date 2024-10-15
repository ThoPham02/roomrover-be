package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf
	DataSource          string
	CreateBillByTimeJob struct {
		Time      string
		BeforeDay int
	}
}
