package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/set/", func(w http.ResponseWriter, r *http.Request) {
		//name value path domain max-age
		cookie := http.Cookie{
			Name:     "tag",
			Value:    "1",
			Path:     "/",
			MaxAge:   60,
			Expires:  time.Now().Add(time.Second * 60),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		// w.Header().Set("Set-Cookie", cookie.String())
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		//name value path domain max-age
		cookie, err := r.Cookie("tag")
		fmt.Println(cookie, err)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":9998", nil)
}
