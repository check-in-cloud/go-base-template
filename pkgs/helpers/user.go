package helpers

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const HeaderUserId = "x-user-id"

func UserIDFromContext(ctx context.Context) (string, error) {
	if mData, ok := metadata.FromIncomingContext(ctx); ok {
		if val, ok := mData[HeaderUserId]; ok {
			return val[0], nil
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String(), nil
	}

	return "unknown", nil
}
