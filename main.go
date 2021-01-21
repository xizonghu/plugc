package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"plugc/compile"
	"plugc/module/logutils"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&logutils.LogFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func main() {
	prepare := compile.PlugcPreparerNew("test.plc")
	prepare.Run()
}