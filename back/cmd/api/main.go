package main

import (
	"github.com/mkmuniz/FinTrack/db"
	"github.com/mkmuniz/FinTrack/config"
	"github.com/mkmuniz/FinTrack/internal/server"
)

func main() {
	config.LoadEnv()
	db.Connect()
	server.Start()
}