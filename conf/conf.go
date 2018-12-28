package conf

import (
	"fmt"
	"log"
	"sync"

	"github.com/xujintao/agollo"
)

var (
	Config config
	mu     sync.RWMutex
)

func init() {
	// agollo.Start()
	err := agollo.StartWithConfFile("meta-config.json")
	if err != nil {
		log.Fatal(err)
	}

	var c config
	if err := agollo.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	Config.Set(&c)

	// 热更新
	agollo.OnConfigChange(func(e *agollo.ChangeEvent) {
		var c config
		if err := agollo.Unmarshal(&c); err != nil {
			log.Println(err)
			return
		}
		fmt.Println(c)
		Config.Set(&c)
	})
}

type config struct {
	// dns 配置
	DNS1 struct {
		ID     string `mapstructure:"id"`
		Token  string `mapstructure:"token"`
		Domain string `mapstructure:"domain"`
	} `mapstructure:"dnspod1"`
	DNS2 struct {
		ID     int    `mapstructure:"id"`
		Token  string `mapstructure:"token"`
		Domain string `mapstructure:"domain"`
	} `mapstructure:"dnspod2.yaml"`

	// DB
	DB struct {
		DSN     string `mapstructure:"dsn"`
		MaxConn string `mapstructure:"max_conn"`
	} `mapstructure:"db"`
}

func (c *config) Set(cc *config) {
	mu.Lock()
	defer mu.Unlock()
	*c = *cc
}

func (c *config) GetDNSID() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.DNS1.ID
}

func (c *config) GetDBMaxConn() string {
	mu.Lock()
	defer mu.RUnlock()
	return c.DB.MaxConn
}
