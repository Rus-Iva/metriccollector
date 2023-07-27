package server

import (
	customlog "github.com/Rus-Iva/metriccollector/internal/logger"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"os"
)

type Server struct {
	storage        *storage.MemStorage
	executablePath string
	Logger         *customlog.Logger
}

func NewServer() *Server {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	logger := customlog.NewLogger()

	s := Server{
		storage:        storage.NewMemStorage(),
		executablePath: ex,
		Logger:         logger,
	}
	return &s
}
