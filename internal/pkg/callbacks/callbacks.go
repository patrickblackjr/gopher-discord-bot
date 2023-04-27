package callbacks

import (
	"golang.org/x/sync/singleflight"
)

type OperationsGateway interface {
	Process(*operations.Request) <-chan singleflight.Result
}

type Handler struct {
	BotName string
}
