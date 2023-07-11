package v2

import (
	"net/http"

	"github.com/traPtitech/go-traq"
)

var UserIdToUserName, UserNameToUserId = UserCache()

func UserCache() (UserIdToUserName, UserNameToUserId map[string]string) {
	User, _, err := GetUsers()
	if err != nil {
		panic(err)
	}

	UserIdToUserName = map[string]string{}
	for _, user := range User {
		UserIdToUserName[user.Id] = user.Name
	}
	UserNameToUserId = map[string]string{}
	for _, user := range User {
		UserNameToUserId[user.Name] = user.Id
	}
	return
}

func GetUsers() ([]traq.User, *http.Response, error) {
	return GetClient(Bot).UserApi.GetUsers(GetContext(Bot)).IncludeSuspended(true).Execute()
}

func UserIdToUserNameFunc(UserId string) string {
	return UserIdToUserName[UserId]
}

func UserNameToUserIdFunc(UserName string) string {
	return UserNameToUserId[UserName]
}
