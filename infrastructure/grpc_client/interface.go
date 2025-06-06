package grpcclient

import (
	"context"

	mpb "github.com/starrysilk/immortal/infrastructure/grpc_client/gen"
)

type IClient interface {
	RegisterService(ctx context.Context,
		port, region string,
	) (*mpb.RegisterServiceResponse, error)
	GetParameters(ctx context.Context) (*mpb.GetParametersResponse, error)
	AddLog(ctx context.Context, msg, stack string) (*mpb.AddLogResponse, error)
	SendReport(ctx context.Context, eid string) (*mpb.SendReportResponse, error)
	SetID(id string)
}
