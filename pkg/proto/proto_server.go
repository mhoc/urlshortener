package proto

import context "context"

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) HealthCheck(context.Context, *HealthChecpReq) (*HealthCheckResp, error) {
	return &HealthCheckResp{
		Ok: true,
	}, nil
}
