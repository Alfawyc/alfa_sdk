package types

import (
	"encoding/json"
	"time"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

func (t *Token) UnmarshalJSON(data []byte) error {
	type xToken Token
	x := &xToken{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.Expiry = time.Now().Add(time.Duration(x.ExpiresIn) * time.Second)
	*t = Token(*x)
	return nil
}

func (t *Token) ExpiryDuration() time.Duration {
	return t.Expiry.Sub(time.Now())
}
