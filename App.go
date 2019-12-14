package main

import (
    "App/database"
    "App/http/processes"
    "App/init/Config"
    "App/init/Logger"
    "App/router"
    "github.com/AboutGoods/go-utils/log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func init() {
    Config.Load()
    Logger.Load()
    database.NewConnection(os.Getenv("DATABASE_URL"))
    processes.PreloadLangData()
}

func main() {

    gin.SetMode(gin.ReleaseMode)

    engine := gin.New()

    engine.Use(gin.Recovery())

    router.DeclareRoutes(engine)

    address := os.Getenv("ADDRESS")
    port := os.Getenv("PORT")

    logrus.WithFields(logrus.Fields{
        "address": address,
        "port": port,
    }).Info("server starting")

    err := engine.Run(address +":"+ port)
    log.PanicOnError(err)
}
