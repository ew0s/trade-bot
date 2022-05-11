package resource

import (
	"io"

	log "github.com/sirupsen/logrus"
)

func Close(logger log.FieldLogger, resource io.Closer) {
	logger.Debugf("Close %T", resource)

	if err := resource.Close(); err != nil {
		logger.WithError(err).Errorf("Can't close %T", resource)
	}
}
