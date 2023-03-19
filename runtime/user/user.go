//go:build osusergo
// +build osusergo

package user

import "os/user"

// CGO is disabled by default, if you need to use CGO,
// please import the github.com/3JoB/ulib/runtime/cuser package instead
// of the github.com/3JoB/ulib/runtime/user package
func UserCurrent() (*user.User, error) {
	return user.Current()
}
