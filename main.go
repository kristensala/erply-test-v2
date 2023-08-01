package main

import "github.com/kristensala/erply-test-v2/config"


func main() {
    config := config.Init()
    router := InitRouter(config)
    router.Run(":5123")
}
