package grpcclient

import (
	"context"

	mpb "github.com/starrysilk/immortal/infrastructure/grpc_client/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	RegistryService   mpb.ServiceRegistryClient
	ParametersService mpb.ParametersClient
	LogService        mpb.LogClient
	ReportService     mpb.ReportClient
	id                string
	config            Config
	conn              *grpc.ClientConn
}

func New(endpoint string, cfg Config) (IClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		RegistryService:   mpb.NewServiceRegistryClient(conn),
		ParametersService: mpb.NewParametersClient(conn),
		LogService:        mpb.NewLogClient(conn),
		config:            cfg,
		conn:              conn,
	}, nil
}

func (c *Client) SetID(id string) {
	c.id = id
}

func (c *Client) RegisterService(ctx context.Context,
	port, region string,
) (*mpb.RegisterServiceResponse, error) {
	return c.RegistryService.RegisterService(ctx, &mpb.RegisterServiceRequest{
		Type:                   mpb.ServiceTypeEnum_RELAY,
		Port:                   port,
		HeartbeatDurationInSec: c.config.Heartbeat,
		Region:                 region,
	})
}

func (c *Client) GetParameters(ctx context.Context) (*mpb.GetParametersResponse, error) {
	md := metadata.New(map[string]string{"x-identifier": c.id})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.ParametersService.GetParameters(ctx, &mpb.GetParametersRequest{})
}

func (c *Client) AddLog(ctx context.Context, msg, stack string) (*mpb.AddLogResponse, error) {
	return c.LogService.AddLog(ctx, &mpb.AddLogRequest{
		Message: msg,
		Stack:   stack,
	})
}

func (c *Client) SendReport(ctx context.Context, eid string) (*mpb.SendReportResponse, error) {
	return c.ReportService.SendReport(ctx, &mpb.SendReportRequest{
		EventId: eid,
	})
}
