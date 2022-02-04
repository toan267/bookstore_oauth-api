package access_token

import (
	"fmt"
	"github.com/toan267/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/toan267/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
	grantTypePassword	= "password"
	grantTypeClientCredentials	="client_credentials"
)

type AccessTokenRequest struct {
	GrantType	string	`json:"grant_type"`
	Scope 	string	`json:"scope"`
	// Used for password grant_type
	UserName	string	`json:"username"`
	Password 	string	`json:"password"`
	//Used for client credentials grant type
	ClientId 	string	`json:"client_id"`
	ClientSecret string `json:"client_secret"`

	//ClientId 	int64 `json:"client_id"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break;
	case grantTypeClientCredentials:
		break;
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")

	}
	//TODO: validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string	`json:"access_token"`
	UserId 		int64 	`json:"user_id"`
	ClientId 	int64 	`json:"client_id,omitempty"`
	Expires		int64	`json:"expires"`
}


func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	//now := time.Now().UTC()
	//expirationTime := time.Unix(at.Expires,0)
	//fmt.Println(expirationTime)
	//return expirationTime.Before(now)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
// web frontend - client-id: 123
// android app - client-id: 234

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}