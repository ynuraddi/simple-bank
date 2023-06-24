package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if useragent := md.Get(grpcGatewayUserAgentHeader); len(useragent) > 0 {
			mtdt.UserAgent = useragent[0]
		}

		if useragent := md.Get(userAgentHeader); len(useragent) > 0 {
			mtdt.UserAgent = useragent[0]
		}

		if clientIP := md.Get(xForwardedForHeader); len(clientIP) > 0 {
			mtdt.ClientIP = clientIP[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
