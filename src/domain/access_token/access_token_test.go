package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	//if expirationTime != 24 {
	//	t.Error("expiration time should be 24 hours")
	//}
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	//if at == nil {
	//	t.Error("Brand new access token should not be nil")
	//}
	assert.False(t, at.isExpired(), "brand new access token should not be expired")
	//if at.isExpired() {
	//	t.Error("brand new access token should not be expired")
	//}
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	//if at.AccessToken != "" {
	//	t.Error("new access token should not have defined access token id")
	//}
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
	//if at.UserId != 0 {
	//	t.Error("new access token should not have an associated user id")
	//}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	//if !at.isExpired() {
	//	t.Error("empty access token should be expired by default")
	//}
	assert.True(t, at.isExpired(), "empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()
	if at.isExpired() {
		t.Error("access token creating three hours from now should NOT be expired")
	}
}