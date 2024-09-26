package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/benosborntech/feedme/apigw/oauth"
	"github.com/benosborntech/feedme/apigw/types"
)

func InjectUserMiddleware(next func(types.Token) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "no authorization header", http.StatusUnauthorized)
			return
		}

		split := strings.Split(authorization, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusBadRequest)
			return
		}

		// ***** We are currently not signing this token, it is currently just JSON

		tokenRaw := split[1]
		var token types.Token
		err := json.Unmarshal([]byte(tokenRaw), &token)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to unmarshal token, err=%v", err), http.StatusBadRequest)
			return
		}

		var oauthObj oauth.OAuth
		switch token.TokenType {
		case types.GoogleType:
			oauthObj = oauth.NewOAuthGoogle()
		default:
			http.Error(w, fmt.Sprintf("unsupported token type, type=%v", token.TokenType), http.StatusBadRequest)
			return
		}

		// **** Now we need to attempt to decode the token...
		// **** We need to extra the details from the header, check the expiry and maybe refresh, then continue forward

		// Execute the handler
		next(token)(w, r)
	}
}
