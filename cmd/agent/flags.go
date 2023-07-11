package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"strconv"
)

var flagBaseEndpointAddr string
var flagReportInterval int64
var flagPollInterval int64

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	flag.StringVarP(&flagBaseEndpointAddr, "addr", "a", "localhost:8080", "address and port to run agent")
	flag.Int64VarP(&flagReportInterval, "repint", "r", 10, "metrics report interval")
	flag.Int64VarP(&flagPollInterval, "polint", "p", 2, "metrics poll interval")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		flagBaseEndpointAddr = envRunAddr
	}
	if envRepInt := os.Getenv("REPORT_INTERVAL"); envRepInt != "" {
		parsedEnvRepInt, err := strconv.ParseInt(envRepInt, 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		flagReportInterval = parsedEnvRepInt
	}
	if envPolInt := os.Getenv("POLL_INTERVAL"); envPolInt != "" {
		parsedEnvPolInt, err := strconv.ParseInt(envPolInt, 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		flagPollInterval = parsedEnvPolInt
	}
}
