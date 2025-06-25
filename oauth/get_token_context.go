package oauth

import (
	"context"
	"sync"
	"time"

	"golang.org/x/oauth2"

	"github.com/sadhasiva1984/openapi"
	"github.com/sadhasiva1984/openapi/models"
	"github.com/sadhasiva1984/openapi/nrf/AccessToken"
)

var tokenMap sync.Map
var clientMap sync.Map

func GetTokenCtx(
	nfType, targetNF models.NrfNfManagementNfType,
	nfId, nrfUri, scope string,
	requesterPlmnId, targetPlmnId *models.PlmnId,
) (context.Context, *models.ProblemDetails, error) {
	tok, pd, err := sendAccTokenReq(nfType, targetNF, nfId, nrfUri, scope, requesterPlmnId, targetPlmnId)
	if err != nil {
		return nil, pd, err
	}
	return context.WithValue(context.Background(),
		openapi.ContextOAuth2, tok), pd, nil
}

func sendAccTokenReq(
	nfType, targetNF models.NrfNfManagementNfType,
	nfId, nrfUri, scope string,
	requesterPlmnId, targetPlmnId *models.PlmnId,
) (oauth2.TokenSource, *models.ProblemDetails, error) {
	var client *AccessToken.APIClient

	configuration := AccessToken.NewConfiguration()
	configuration.SetBasePath(nrfUri)

	if val, ok := clientMap.Load(configuration); ok {
		client = val.(*AccessToken.APIClient)
	} else {
		client = AccessToken.NewAPIClient(configuration)
		clientMap.Store(configuration, client)
	}

	var tok models.NrfAccessTokenAccessTokenRsp
	if val, ok := tokenMap.Load(scope); ok {
		tok = val.(models.NrfAccessTokenAccessTokenRsp)
		if int32(time.Now().Unix()) < tok.ExpiresIn {
			token := &oauth2.Token{
				AccessToken: tok.AccessToken,
				TokenType:   tok.TokenType,
				Expiry:      time.Unix(int64(tok.ExpiresIn), 0),
			}
			return oauth2.StaticTokenSource(token), nil, nil
		}
	}

	req := &AccessToken.AccessTokenRequestRequest{}
	req.SetGrantType("client_credentials")
	req.SetNfInstanceId(nfId)
	req.SetNfType(nfType)
	req.SetTargetNfType(targetNF)
	req.SetScope(scope)

	// Set PLMN IDs if provided
	if requesterPlmnId != nil {
		req.SetRequesterPlmn(*requesterPlmnId)
	}
	if targetPlmnId != nil {
		req.SetTargetPlmn(*targetPlmnId)
	}

	res, err := client.AccessTokenRequestApi.AccessTokenRequest(
		context.Background(), req)

	if err == nil {
		tokenMap.Store(scope, res.NrfAccessTokenAccessTokenRsp)
		token := &oauth2.Token{
			AccessToken: res.NrfAccessTokenAccessTokenRsp.AccessToken,
			TokenType:   res.NrfAccessTokenAccessTokenRsp.TokenType,
			Expiry:      time.Unix(int64(res.NrfAccessTokenAccessTokenRsp.ExpiresIn), 0),
		}
		return oauth2.StaticTokenSource(token), nil, nil
	} else {
		return nil, nil, openapi.ReportError("server no response")
	}
}
