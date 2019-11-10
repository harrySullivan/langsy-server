package Logger

import (
    log "github.com/sirupsen/logrus"
    "os"
)

func Load() {
    log.SetOutput(os.Stdout)
    env := os.Getenv("ENV")
    log.SetLevel(log.DebugLevel)
    if env == "prod" {
        log.SetFormatter(&log.JSONFormatter{})
        log.SetLevel(log.InfoLevel)
    }

}
