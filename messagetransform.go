package messagetransform

import (
	"github.com/kr/pretty"
)

func Transform(messageRef *map[string]interface{}, transform map[string]interface{}) error {
	message := *messageRef
	message["x"] = "y"
	return nil
}
func no() {
	pretty.Println("no")
}
