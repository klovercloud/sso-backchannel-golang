package dto

type BackChannelAuthDto struct {
	Code        string `json:"code"`
	RedirectUri string `json:"redirectUri"`
}

type FrontChannelInfoDto struct {
	AuthorizeURI string `json:"authorizeURI"`
	RedirectURI  string `json:"redirectURI"`
	ClientID     string `json:"clientID"`
	Scope        string `json:"scope"`
}
