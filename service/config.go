package service

import (
	"go.uber.org/zap"

	"golang-program-structure/common/middleware"
)

type Config struct {
	Env         string `required:"true"`
	ServiceName string `required:"true"`
	ServicePort string `required:"true"`
	Version     string `required:"true"`

	Mw middleware.MiddlewareConfig `required:"true"`

	Logger *zap.SugaredLogger `ignored:"true"`
}
