package dto

type T struct {
	Exp            int    `json:"exp"`
	Iat            int    `json:"iat"`
	AuthTime       int    `json:"auth_time"`
	Jti            string `json:"jti"`
	Iss            string `json:"iss"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Typ            string `json:"typ"`
	Azp            string `json:"azp"`
	Nonce          string `json:"nonce"`
	SessionState   string `json:"session_state"`
	Acr            string `json:"acr"`
	ResourceAccess struct {
		ClientAppTwo struct {
			Roles []string `json:"roles"`
		} `json:"client-app-two"`
		ClientAppOne struct {
			Roles []string `json:"roles"`
		} `json:"client-app-one"`
	} `json:"resource_access"`
	Scope             string `json:"scope"`
	Sid               string `json:"sid"`
	EmailVerified     bool   `json:"email_verified"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}
