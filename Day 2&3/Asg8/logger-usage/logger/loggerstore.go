package logger

import "log"

type LoggerStore struct {
	log *log.Logger
}

func New() *LoggerStore {
	return &LoggerStore{
		log: log.New(log.Writer(), "CustomLogger: ", log.LstdFlags),
	}
}