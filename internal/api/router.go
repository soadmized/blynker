package api

import "net/http"

func (a *API) Route(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/get_data" {
		if req.Method == http.MethodPost {
			a.Get(w, req)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			makeResponse(w, "WRONG METHOD, USE POST")
		}
	}

	if req.URL.Path == "/set_data" {
		if req.Method == http.MethodPost {
			a.Save(w, req)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			makeResponse(w, "WRONG METHOD, USE POST")
		}
	}

	if req.URL.Path == "/" {
		if req.Method == http.MethodGet {
			a.DisplayValues(w, req)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			makeResponse(w, "WRONG METHOD, USE GET")
		}
	}
}
