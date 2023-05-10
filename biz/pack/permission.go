package pack

import (
	"Refined_service/biz/dal/sqlite"
	"Refined_service/biz/model/user"
)

func CheckAuthValid(username, password string) bool {
	u, _, err := sqlite.QueryBasic[user.User]("username = ?", username)
	if err != nil {
		return false
	}
	if HashSHA256(password) != u[0].Password {
		return false
	}
	return true
}
