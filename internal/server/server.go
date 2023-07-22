package server

import (
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"os"
)

type Server struct {
	storage        *storage.MemStorage
	executablePath string
}

func NewServer() *Server {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	s := Server{
		storage:        storage.NewMemStorage(),
		executablePath: ex,
	}
	return &s
}
