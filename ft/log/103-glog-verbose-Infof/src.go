// Issue 89
// Passing tainted data into glog.Verbose.Infof can
// result in log injection.
package main

import (
	"net/http"

	"github.com/golang/glog"
)

type User struct {
	username string
	password string
}

func newUser(name string, pass string) User {
	return User{
		username: name,
		password: pass,
	}
}

func Log(userArr [10]*User, i int) {
	glog.V(1).Infof("Failed to log user: %s", userArr[i].username)
}

func handler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	var userArr [10]*User
	user := newUser(username, password)
	userArr[0] = &user
	w.Write([]byte("Hello, world!"))
	Log(userArr, 0)

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
