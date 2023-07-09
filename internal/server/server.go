package server

import (
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"sync"
)

type Server struct {
	sync.Mutex
	storage *storage.MemStorage
}

func NewServer() *Server {
	s := Server{
		storage: storage.NewMemStorage(),
	}
	return &s
}
