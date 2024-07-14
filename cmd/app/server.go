package main

import (
	"GOAuth/config"
	dao "GOAuth/internal/dao/users"
	"GOAuth/internal/must"
	"GOAuth/internal/services"
	"GOAuth/migration"
	"context"
	"fmt"
	"log"
)

func main() {
	var ctx = context.TODO()
	cfg := config.ReadConfigAndArg()
	fmt.Println("Environment: ", cfg.Env, "Port: ", cfg.Port)
	db := must.ConnectDb(cfg.Db)
	fmt.Println("Db connected successfully")
	err := migration.Migration(db)
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}

	if err := migration.AutoSeedingData(db); err != nil {
	}
	//dao
	useDao := dao.NewUser(db)
	must.NewServer(ctx, cfg, nil, services.NewGokitPublicService(cfg, useDao))
}
