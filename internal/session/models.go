package session

const AuthCookieName = "lunchbunch_auth"
const SettingsCookieName = "lunchbunch_settings"
const AuthenticationStoreKey = "authenticated"

type Data struct {
	IsAuthenticated bool
	Username        string
}
