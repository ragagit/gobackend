package handlers

import (
	"net/http"

	"github.com/ragagit/GoBackEnd/section6/gopherfaceauth/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
