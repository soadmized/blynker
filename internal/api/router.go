package api

import (
	"net/http"
)

func (a *API) Route(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/get_data" {
		if a.checkAuth(req) {
			if req.Method == http.MethodPost {
				a.GetData(w, req)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				makeResponse(w, "WRONG METHOD, USE POST")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			makeResponse(w, "AUTH REQUIRED!")
		}
	}

	if req.URL.Path == "/set_data" {
		if a.checkAuth(req) {
			if req.Method == http.MethodPost {
				a.SaveData(w, req)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				makeResponse(w, "WRONG METHOD, USE POST")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			makeResponse(w, "AUTH REQUIRED!")
		}
	}

	if req.URL.Path == "/" {
		if req.Method == http.MethodGet {
			a.GetStatus(w, req)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			makeResponse(w, "WRONG METHOD, USE GET")
		}
	}
}

func (a *API) checkAuth(req *http.Request) bool {
	user, pass, _ := req.BasicAuth()
	if (user == a.conf.User) && (pass == a.conf.Pass) {
		return true
	}

	return false
}
