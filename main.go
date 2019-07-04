package main

import (
    "blogkit/app"
    "blogkit/config"
    "fmt"
)

func main() {
    config.LoadConfig()
    fmt.Println("Blog Kit version " + config.Cfg.Version)
    app.Run()
}
