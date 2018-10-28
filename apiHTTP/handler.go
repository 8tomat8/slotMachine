package apiHTTP

import (
	"go.uber.org/zap"
)

type Handler struct {
	Log       *zap.SugaredLogger
	JWTSecret []byte
}
