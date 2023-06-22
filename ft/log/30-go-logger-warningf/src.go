// Issue 89
// Passing tainted data into logger.Warningf can
// result in log injection.
package testdata

import (
	"net/http"

	"os"

	"github.com/apsdehal/go-logger"
)

func handler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	logger, err := logger.New("MyLogger", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	logger.Warningf("Failed to log in %s", username)
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
