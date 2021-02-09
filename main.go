package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"plugc/compiler"
	"plugc/module/logutils"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&logutils.LogFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func main() {
	prepare := compiler.PlugcPreparerNew("test.plc")
	prepare.Run()
}