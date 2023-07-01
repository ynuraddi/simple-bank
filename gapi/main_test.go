package gapi

import (
	"testing"
	"time"

	db "simple-bank/db/sqlc"
	"simple-bank/util"
	"simple-bank/worker"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

// func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, duration time.Duration) context.Context {
// 	accessToken, _, err := tokenMaker.CreateToken(username, duration)
// 	require.NoError(t, err)

// 	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
// 	md := metadata.MD{
// 		authorizationHeader: []string{
// 			bearerToken,
// 		},
// 	}

// 	return metadata.NewIncomingContext(context.Background(), md)
// }
