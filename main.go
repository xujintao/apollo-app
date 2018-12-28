package main

import (
	"fmt"

	"github.com/xujintao/apollo-app/conf"
)

// go mod init github.com/xujintao/apollo-app
// go mod tidy -v

func main() {
	fmt.Println(conf.Config.GetDBMaxConn())
	select {}
}
