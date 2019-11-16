package util

import (
	"github.com/devfeel/dotweb"
	"strconv"
)

func NewSimpleCROS() dotweb.Middleware {
	option := NewConfig().UseDefault().SetAllowCredentials(true).SetMaxAge(3600)
	mid := Middleware(option)
	return mid
}

//CROS配置
type Config struct {
	enabledCROS      bool
	allowedOrigins   string
	allowedMethods   string
	allowedHeaders   string
	allowCredentials bool
	exposeHeaders    string
	allowedP3P       string
	maxAge           int
}

func (c *Config) UseDefault() *Config {
	c.enabledCROS = true
	c.allowedOrigins = "*"
	c.allowCredentials = true
	c.allowedMethods = "GET,HEAD,POST,PUT,DELETE,PATCH"
	c.allowedHeaders = "Content-Type,Accept,Total,Cookie"
	//c.allowedHeaders = "Content-Type,x-requested-with,multipart/form-assets12,text/plain,application/x-www-form-urlencoded"
	c.exposeHeaders = "X-Pagination-Current-Page,Content-Type,Total,total,Set-Cookie"
	c.allowedP3P = "CP=\"CURa ADMa DEVa PSAo PSDo OUR BUS UNI PUR INT DEM STA PRE COM NAV OTC NOI DSP COR\""
	return c
}

func (c *Config) Enabled() *Config {
	c.enabledCROS = true
	return c
}

func (c *Config) SetOrigin(origins string) *Config {
	c.allowedOrigins = origins
	return c
}

func (c *Config) SetMethod(methods string) *Config {
	c.allowedMethods = methods
	return c
}

func (c *Config) SetHeader(headers string) *Config {
	c.allowedHeaders = headers
	return c
}

func (c *Config) SetExposeHeaders(headers string) *Config {
	c.exposeHeaders = headers
	return c
}

func (c *Config) SetAllowCredentials(flag bool) *Config {
	c.allowCredentials = flag
	return c
}

func (c *Config) SetMaxAge(maxAge int) *Config {
	c.maxAge = maxAge
	return c
}

func (c *Config) SetP3P(p3p string) *Config {
	c.allowedP3P = p3p
	return c
}

func NewConfig() *Config {
	return &Config{}
}

//jwt中间件
type CORSMiddleware struct {
	dotweb.BaseMiddlware
	config *Config
}

func (m *CORSMiddleware) Handle(ctx dotweb.Context) error {
	if m.config.enabledCROS {

		if m.config.allowedOrigins == "*" {
			ctx.Response().SetHeader(dotweb.HeaderAccessControlAllowOrigin, ctx.Request().QueryHeader("Origin"))
		} else {
			ctx.Response().SetHeader(dotweb.HeaderAccessControlAllowOrigin, m.config.allowedOrigins)
		}
		ctx.Response().SetHeader(dotweb.HeaderAccessControlAllowMethods, m.config.allowedMethods)
		ctx.Response().SetHeader(dotweb.HeaderAccessControlAllowHeaders, m.config.allowedHeaders)
		ctx.Response().SetHeader(dotweb.HeaderAccessControlExposeHeaders, m.config.exposeHeaders)
		ctx.Response().SetHeader(dotweb.HeaderAccessControlAllowCredentials, strconv.FormatBool(m.config.allowCredentials))
		ctx.Response().SetHeader(dotweb.HeaderAccessControlMaxAge, strconv.Itoa(m.config.maxAge))
		ctx.Response().SetHeader(dotweb.HeaderP3P, m.config.allowedP3P)
		//ctx.Response().SetHeader("Connection","keep-alive")
		if ctx.Request().Method == "OPTIONS" {
			ctx.Response().Write(203, nil)
			//ctx.End()
			return nil
		}
	}
	return m.Next(ctx)
}

// Middleware create new CORS Middleware
func Middleware(config *Config) *CORSMiddleware {
	return &CORSMiddleware{config: config}
}

// DefaultMiddleware create new CORS Middleware with default config
func DefaultMiddleware() *CORSMiddleware {
	option := NewConfig().UseDefault()
	return &CORSMiddleware{config: option}
}
