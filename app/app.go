package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore(".sessions/", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
