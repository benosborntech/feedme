package oauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/benosborntech/feedme/apigw/consts"
	"github.com/benosborntech/feedme/apigw/types"
	"github.com/benosborntech/feedme/common/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthGoogle struct {
	client *redis.Client
	config *oauth2.Config
}

type googleProfile struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func NewOAuthGoogle(client *redis.Client, config *types.OAuthConfig, baseURL string) *OAuthGoogle {
	out := &OAuthGoogle{
		client: client,
		config: &oauth2.Config{
			ClientID:     config.ClientId,
			ClientSecret: config.ClientSecret,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}

	out.config.RedirectURL = fmt.Sprintf("%s/%s", baseURL, out.GetServiceType())

	return out
}

func (o *OAuthGoogle) GetServiceType() types.ServiceType {
	return types.GoogleType
}

func (o *OAuthGoogle) GetEndpointPath() string {
	return fmt.Sprintf("/auth/login/%s", o.GetServiceType())
}

func (o *OAuthGoogle) GetEndpoint(ctx context.Context) (session string, endpoint string, err error) {
	state, err := utils.GenerateRand(consts.STATE_LENGTH)
	if err != nil {
		return "", "", err
	}

	session, err = utils.GenerateRand(consts.SESSION_LENGTH)
	if err != nil {
		return "", "", err
	}

	res, err := o.client.SetNX(ctx, session, state, consts.SESSION_EXPIRY).Result()
	if err != nil {
		return "", "", err
	}
	if !res {
		return "", "", errors.New("failed to set session key")
	}

	return session, o.config.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

func (o *OAuthGoogle) GetCallbackPath() string {
	return fmt.Sprintf("/auth/callback/%s", o.GetServiceType())
}

func (o *OAuthGoogle) ExchangeToken(ctx context.Context, code string, state string, session string) (*oauth2.Token, error) {
	sessionState, err := o.client.Get(ctx, session).Result()
	if err != nil {
		return nil, err
	}
	defer o.client.Del(ctx, session)

	if sessionState != state {
		return nil, errors.New("state does not match session state")
	}

	t, err := o.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (o *OAuthGoogle) GetUserInfo(ctx context.Context, accessToken string) (*types.UserInfo, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profile googleProfile
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}

	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	// Decode the payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}

	// Get the sub claim
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("could not get sub")
	}

	userInfo := &types.UserInfo{
		Sub:   sub,
		Email: profile.Email,
		Name:  profile.Name,
	}

	return userInfo, nil
}

func (o *OAuthGoogle) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}

	tokenSource := o.config.TokenSource(ctx, token)

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("unable to refresh token: %w", err)
	}

	return newToken, nil
}
