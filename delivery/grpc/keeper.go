package grpc

import (
	mpb "github.com/starrysilk/immortal/delivery/grpc/gen"
	"github.com/starrysilk/immortal/delivery/websocket"
	"github.com/starrysilk/immortal/pkg/utils"
	"github.com/starrysilk/immortal/repository"
)

type ParametersKeeper struct {
	Handler         *repository.Config
	WebsocketServer *websocket.Config
}

func (keeper *ParametersKeeper) LoadParameters(params *mpb.UpdateParametersRequest) error {
	url, err := utils.ParseURL(params.Url)
	if err != nil {
		return err
	}

	keeper.WebsocketServer.SetURL(url)

	keeper.WebsocketServer.SetLimitation(
		&websocket.Limitation{
			MaxMessageLength:    params.Limitations.MaxMessageLength,
			MaxSubscriptions:    params.Limitations.MaxSubscriptions,
			MaxSubidLength:      params.Limitations.MaxSubidLength,
			MinPowDifficulty:    params.Limitations.MinPowDifficulty,
			AuthRequired:        params.Limitations.AuthRequired,
			PaymentRequired:     params.Limitations.PaymentRequired,
			RestrictedWrites:    params.Limitations.RestrictedWrites,
			MaxEventTags:        params.Limitations.MaxEventTags,
			MaxContentLength:    params.Limitations.MaxContentLength,
			CreatedAtLowerLimit: params.Limitations.CreatedAtLowerLimit,
			CreatedAtUpperLimit: params.Limitations.CreatedAtUpperLimit,
		})

	keeper.Handler.SetMaxQueryLimit(params.Limitations.MaxQueryLimit)
	keeper.Handler.SetDefaultQueryLimit(params.Limitations.DefaultQueryLimit)

	return nil
}
