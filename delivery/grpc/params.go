package grpc

import (
	"context"

	mpb "github.com/starrysilk/immortal/delivery/grpc/gen"
)

type paramsServer struct {
	*Server
}

func newParamsServer(s *Server) *paramsServer {
	return &paramsServer{
		Server: s,
	}
}

func (s *paramsServer) UpdateParameters(_ context.Context,
	newParams *mpb.UpdateParametersRequest,
) (*mpb.UpdateParametersResponse, error) {
	err := s.keeper.LoadParameters(newParams)

	return &mpb.UpdateParametersResponse{Success: err == nil}, err
}
