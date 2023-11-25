package session

const CookieName = "lunchbunch_auth"
const AuthenticationStoreKey = "authenticated"

type Data struct {
	Authenticated bool
	Username      string
}
