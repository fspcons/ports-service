package file

import (
	"context"
	"github.com/fspcons/ports-service/src/config"
	"github.com/fspcons/ports-service/src/domain"
	"go.uber.org/zap"
)

// Gateway allows checking on json ports file.
//
//go:generate moq -out gateway_mock.go . Gateway:GatewayMock
type Gateway interface {
	// CheckOnFile if the port record exists on the json file
	CheckOnFile(ctx context.Context, port *domain.Port) error
}

// NewFileGateway builds a new file gateway.
func NewFileGateway(cfg config.Data, logger *zap.Logger) Gateway {
	return &fileChecker{logger: logger, portFilePath: cfg.PortsFilePath}
}
