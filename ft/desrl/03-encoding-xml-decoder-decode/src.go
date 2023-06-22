// Issue 89
// Passing tainted data into encoding.xml.Decoder.Decode can
// result in Deserialization attack.

package testdata

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer file.Close()

	var person Person
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)

	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
