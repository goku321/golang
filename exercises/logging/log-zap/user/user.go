package user

import (
	"go.uber.org/zap"
)

func Name() {
	zap.S().Info("Logging from user package")
}