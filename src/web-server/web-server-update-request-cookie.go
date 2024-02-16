package web_server

import "net/http"

func updateRequestCookie(r *http.Request, cookie *http.Cookie) {
	cookieHeader := ""
	for _, c := range r.Cookies() {
		if c.Name == cookie.Name {
			continue
		}
		cookieHeader += c.String() + "; "
	}

	if cookie.Value != "" {
		cookieHeader += cookie.String() + ";"
	}

	r.Header.Set("Cookie", cookieHeader)
}
