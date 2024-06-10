package main

import (
	"GOAuth/config"
	"GOAuth/internal/must"
	"fmt"
)

func main() {
	cfg := config.ReadConfigAndArg()
	fmt.Println("Environment: ", cfg.Env, "Port: ", cfg.Port)
	db := must.ConnectDb(cfg.Db)
	fmt.Println("Db connected successfully", db)
}
