package main

import (
	flag "github.com/spf13/pflag"
)

var flagBseEndpointAddr string
var flagReportInterval int
var flagPollInterval int

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	flag.StringVarP(&flagBseEndpointAddr, "", "a", ":8080", "address and port to run agent")
	flag.IntVarP(&flagReportInterval, "repint", "r", 10, "metrics report interval")
	flag.IntVarP(&flagPollInterval, "polint", "p", 2, "metrics poll interval")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
