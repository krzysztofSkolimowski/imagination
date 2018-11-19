package modules

import (
	"github.com/krzysztofSkolimowski/imagination/pkg/app/image"
	"github.com/sirupsen/logrus"
)

type Services struct {
	ImageService *image.Service
	Logger       *logrus.Logger
}

func NewServices(svc *image.Service, logger *logrus.Logger) *Services {
	return &Services{ImageService: svc, Logger: logger}
}
