package client

import (
	"context"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)
	if client, err := grpc.Dial(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(c.client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}

	return c, nil
}

func (g *GRPCClient) CreateAuth(name string) (*auth.AuthData, error) {
	now := time.Now()
	expireTime := now.Add(30 * time.Minute)

	a := &auth.AuthData{
		Name:       name,
		CreateDate: now.Unix(),
		ExpireDate: expireTime.Unix(),
	}

	if token, err := g.pasetoMaker.CreateNewToken(*a); err != nil {
		return nil, err
	} else {
		a.Token = token

		if res, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: a}); err != nil {
			return nil, err
		} else {
			return res.Auth, nil
		}
	}

}

func (g *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {
	if res, err := g.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res.V, nil
	}
}
