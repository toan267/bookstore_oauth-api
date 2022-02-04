package access_token

import (
	"github.com/toan267/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string	`json:"access_token"`
	UserId 		int64 	`json:"user_id"`
	ClientId 	int64 	`json:"client_id"`
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

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
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
