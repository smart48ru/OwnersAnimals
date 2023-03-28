package main

import (
	"OwnersAnimals/internal/config"
	"OwnersAnimals/internal/infrastructure/api/routergin"
	"OwnersAnimals/internal/infrastructure/api/server"
	"OwnersAnimals/internal/infrastructure/handlers"
	"OwnersAnimals/internal/infrastructure/repo/animalrepo"
	"OwnersAnimals/internal/infrastructure/repo/ownerrepo"
	"OwnersAnimals/internal/infrastructure/service"
	"OwnersAnimals/internal/infrastructure/store/db"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	dbCfg := config.DB{
		Enable:   false,
		Host:     "127.0.0.1",
		Post:     5432,
		User:     "DBUser",
		Password: "DBPassword",
		DBName:   "owner-animal",
		SSLMode:  false,
	}

	con, err := db.NewPostgresConnect(dbCfg)
	if err != nil {
		log.Fatal().Msgf("connect to DB | %s", err)
	}
	animalStore := db.NewAnimalStore(con)
	ownerStore := db.NewOwnerStore(con)

	animalRepo := animalrepo.NewAnimalRepo(animalStore)
	ownerRepo := ownerrepo.NewOwnerRepo(ownerStore)
	srv := service.NewService(animalRepo, ownerRepo)
	h := handlers.NewHandler(srv)
	r := routergin.NewRouterGin(h)
	apiCfg := config.API{
		APIHost:            "localhost",
		APIPort:            8080,
		APIReadTimeOut:     20,
		APIWriteTimeOut:    30,
		APIReadHeadTimeOut: 30,
	}
	serv := server.NewServer(apiCfg, r)
	serv.Start(srv)

	fmt.Println("Program Start")
	<-ctx.Done()
	fmt.Println("Program Stop")
	serv.Stop()
	cancel()

}
