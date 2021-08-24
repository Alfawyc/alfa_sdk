package types

type RequestMeta struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code,omitempty"`
	RedirectUri  string `json:"redirect_uri,omitempty"`
}

var DefaultPageCount int = 100
