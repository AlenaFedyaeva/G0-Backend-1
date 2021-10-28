package server

import (
	"homework4/db/mydb"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Server struct {
	db     *mydb.DB
	server *sqlx.DB
	log    zerolog.Logger
}

func NewServer(conn string, log zerolog.Logger) (*Server, error) {
	//con, err := sqlx.Connect("postgres", conn)
	return &Server{db: mydb.New(), server: nil, log: log}, nil
}

func (s Server) Close() error {
	return s.server.Close()
}
