package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/simple-event-app/backend/models/repository"
	"github.com/sirupsen/logrus"

	"github.com/simple-event-app/backend/config"
	"github.com/simple-event-app/backend/db"
	"github.com/simple-event-app/backend/module"
	"github.com/simple-event-app/backend/server"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init()
	//Db init
	db := db.Init()

	var err error

	module.Configure.Repository.EventRepository = &repository.EventMysql{DbCon: db}
	module.Configure.Repository.TagRepository = &repository.TagMysql{DbCon: db}

	// server init
	r := server.NewRouter(db)
	if err = r.Run(); err != nil {
		logrus.Errorf("server run error: %s", err)
	}

	defer db.Close()
}
