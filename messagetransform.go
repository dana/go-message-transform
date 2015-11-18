package messagetransform

import (
	"fmt"
	"github.com/kr/pretty"
)

type TransformError struct {
	What string
}

func (e TransformError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

func transformMapMap(messageRef *map[string]interface{}, transform map[string]interface{}) error {
	message := *messageRef
	for key, value := range transform {
		if _, ok := message[key]; !ok {
			message[key] = value
		} else {
			switch value.(type) {
			case string:
				message[key] = value.(string)
			case int:
				message[key] = value.(int)
			case uint:
				message[key] = value.(uint)
			case float64:
				message[key] = value.(float64)
			case map[string]interface{}:
				switch message[key].(type) {
				case map[string]interface{}:
					thing := message[key].(map[string]interface{})
					return transformMapMap(&thing, value.(map[string]interface{}))
				}
			default:
				return TransformError{"unhandled type"}
			}
		}
	}
	return nil
}

func Transform(messageRef *map[string]interface{}, transform map[string]interface{}) error {
	return transformMapMap(messageRef, transform)
}
func no() {
	pretty.Println("no")
}
