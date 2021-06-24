package main

import (
  log "github.com/sirupsen/logrus"  //日志库
)

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
  }).Info("A walrus appears")
}