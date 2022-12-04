package svc

import (
	"github.com/chenkaiwei/crypto-m/cryptom"
	"github.com/chenkaiwei/crypto-m/cryptom/algom"
	"github.com/chenkaiwei/crypto-m/samples/customDemo/internal/config"
	"github.com/chenkaiwei/crypto-m/samples/customDemo/internal/cryptomx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	CryptomRequest  rest.Middleware
	CryptomResponse rest.Middleware
}

const PRIVATE_KEY = "MIICWwIBAAKBgQC9G7U67pfDHfAfUvMkZ2uBPyEbvqYrlr5xEr2zvcoKxXfdfh4+n//ycD2wTE4EkwwAJiVAqaT3s1KilhkMf4RnB5sE1Fj0Aq7n+8tYsmTdK95BIEeSGuO2qZIni5S7EaAZVSQiv8HGbedlteAv2Ja1XRZsnIoe0E9qQBOCAhpB6QIDAQABAoGAV/Kd62VxISY4OWkreP+8CKTicfPNdjIqKY4suX4Hi9DgeRshV8CzmP3IQsiJ9CirCRq0cokzFpvIT6L8zUo0uaRJdaA7cpo2iJLUtoOHc6zdUHdBHBaVEa9ymjEqx2wrIw5kkcy4SBe472DWJDtphgig9S3THww9k4jzn0L6DL0CQQDxsfPRm3nwOj8qaVHBJCYcJ0+rdXL9a7UIYomUBEF6Z+DRlG6iYI5ZXSJZ+Sxhp20QawVk47JlMOGGRM7o2ApbAkEAyEz7/5kkp6IZ3lC4yRrTD86eKjJ9yzSXiOBQiVmW54lluxvK6o4sXP/+hVb9zqiD5AYOwjFbLiCTtSLZgv5wCwJAZSuxPPdQ5p7rG+y0HR3tmfFWpxXlyXDRea4NmtjhM8TR1cjFOtEiJQQYQgNMcaAsxieWPXIWlccNUC/zUIJGawJAKpudu3hjQLmN0SnQtQ7cuO8V3BoTgkd0uKwm1aDWJfinSE8YMh7+NuZJySmBIhXcwIO9XffL0pshcJWyOVhQkwJANTQejTHOgu7bPkGe5Zg+HfHQGMrimffbWckXMTDQpA1y5l5fq7+B7cauSYoNcQZ41Bp38USCzsxkwTn5659iFA=="

func NewServiceContext(c config.Config) *ServiceContext {

	mca := &cryptomx.MyCaesarAlgo{}

	cryptomManager := cryptom.NewStandardCryptomManager(algom.NewCekAlgoRsaBase64(PRIVATE_KEY), mca)
	//⬆️因为不对称加密我手搓不出来，所以还是沿用rsa，仅用content加密策略做演示

	return &ServiceContext{
		Config:          c,
		CryptomRequest:  cryptomManager.RequestHandle,
		CryptomResponse: cryptomManager.ResponseHandle,
	}
}
