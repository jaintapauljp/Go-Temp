// Issue 89
// Passing tainted data into logrus.logger.Warnln can
// result in log injection.
package testdata

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.WarnLevel)
	logrusLogger.Warnln(username)
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
