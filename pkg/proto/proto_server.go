package proto

import (
	"context"
	fmt "fmt"
	"time"

	"gitlab.com/mhoc/urlshortener/pkg/config"
	"gitlab.com/mhoc/urlshortener/pkg/store"
	"gitlab.com/mhoc/urlshortener/pkg/util"
)

type Server struct {
	cfg *config.Config
	st  store.Store
}

func NewServer(cfg *config.Config, st store.Store) Server {
	return Server{
		cfg: cfg,
		st:  st,
	}
}

func (s Server) CreateShortlink(ctx context.Context, req *CreateShortlinkReq) (*CreateShortlinkResp, error) {
	var expireIn time.Duration
	if req.ExpiresInSeconds != nil {
		expireIn = time.Duration(*req.ExpiresInSeconds) * time.Second
	}
	id, err := s.st.Create(req.Url, expireIn)
	if err != nil {
		return nil, fmt.Errorf("error creating short url: %v", err.Error())
	}
	return &CreateShortlinkResp{
		ShortUrl: util.IDToShortlink(s.cfg.RootURL, id),
	}, nil
}

func (s Server) HealthCheck(ctx context.Context, req *HealthChecpReq) (*HealthCheckResp, error) {
	return &HealthCheckResp{
		Ok: true,
	}, nil
}

func (s Server) RemoveShortlink(ctx context.Context, req *RemoveShortlinkReq) (*RemoveShortlinkResp, error) {
	id, err := util.ShortlinkToID(req.ShortUrl)
	if err != nil {
		return nil, fmt.Errorf("provided an invalid short link")
	}
	removed, err := s.st.Remove(id)
	if err != nil {
		return nil, fmt.Errorf("error creating short url: %v", err.Error())
	}
	return &RemoveShortlinkResp{
		Removed: removed,
	}, nil
}
