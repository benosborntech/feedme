package oauth

import (
	"context"

	"github.com/benosborntech/feedme/apigw/types"
	"golang.org/x/oauth2"
)

type OAuth interface {
	GetServiceType() types.ServiceType
	GetEndpointPath() string
	GetEndpoint(context.Context) (string, string, error)
	GetCallbackPath() string
	ExchangeToken(ctx context.Context, code string, state string, session string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, accessToken string) (*types.UserInfo, error)
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
}
