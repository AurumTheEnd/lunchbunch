package session

const CookieName = "lunchbunch_auth"
const AuthenticationStoreKey = "authenticated"

type Data struct {
	IsAuthenticated bool
	Username        string
}
