package authenticate

import (
	"log"
	"strings"

	"github.com/ragagit/GoBackEnd/section6/gopherfaceauth/common"
	"github.com/ragagit/GoBackEnd/section6/gopherfaceauth/common/utility"
)

func VerifyCredentials(e *common.Env, username string, password string) bool {

	u, err := e.DB.GetUser(username)
	if u == nil {
		return false
	}

	if err != nil {
		log.Print(err)
	}

	if strings.ToLower(username) == strings.ToLower(u.Username) && utility.SHA256OfString(password) == u.PasswordHash {
		log.Println("Successful login attempt from user: ", u.Username)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user: ", u.Username)
		return false
	}

}