// Passing non-pointer arg into xml.Decode
// causes runtime error

package testdata

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type MyStruct struct {
}

func main() {
	xmlData := `
	<MyStruct>
		<field1>Hello</field1>
		<field2>42</field2>
	</MyStruct>`

	var result MyStruct

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	err := decoder.Decode(result)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	fmt.Println("Decoded struct:", result)
}

//<<<<<345, 367>>>>>
