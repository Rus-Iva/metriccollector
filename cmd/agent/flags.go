package main

import (
	flag "github.com/spf13/pflag"
)

var flagBseEndpointAddr string
var flagReportInterval int64
var flagPollInterval int64

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	flag.StringVarP(&flagBseEndpointAddr, "addr", "a", "localhost:8080", "address and port to run agent")
	flag.Int64VarP(&flagReportInterval, "repint", "r", 10, "metrics report interval")
	flag.Int64VarP(&flagPollInterval, "polint", "p", 2, "metrics poll interval")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
