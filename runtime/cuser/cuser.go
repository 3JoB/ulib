package cuser

import "os/user"

func UserCurrent() (*user.User, error) {
	return user.Current()
}
