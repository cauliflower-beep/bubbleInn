package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	getCaptcha "getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Handle(ctx context.Context, msg *getCaptcha.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getCaptcha.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
