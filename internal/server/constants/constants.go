package constants

import "fmt"

const IndexPath = "/"
const RegisterFormPath = "/register"
const LoginFormPath = "/login"
const LogoutPath = "/logout"
const NewPollFormPath = "/poll/new"

const pollPath = "/poll/%d"

func PollPath(id uint) string {
	return fmt.Sprintf(pollPath, id)
}
