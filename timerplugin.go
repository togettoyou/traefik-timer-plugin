package timerplugin

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Config 自定义配置
type Config struct {
	Log bool `json:"log,omitempty"`
}

// CreateConfig 提供给 traefik 设置配置
func CreateConfig() *Config {
	return &Config{}
}

// New 提供给 traefik 创建 Timer 插件
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Timer{
		next: next,
		name: name,
		log:  config.Log,
	}, nil
}

type Timer struct {
	next http.Handler
	name string
	log  bool
}

func (t *Timer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()
	t.next.ServeHTTP(rw, req)
	cost := time.Since(start)
	if t.log {
		fmt.Println("请求花费时间：", cost)
	}
}
