package messagetransform

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Transform(message *map[string]interface{}, transform map[string]interface{}) error {
//basic tests from reference implementation:
//https://github.com/dana/perl-Message-Transform/blob/master/t/basic.t
func TestSimplestTransform(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
	}
	transform := map[string]interface{}{
		"x": "y",
	}
	transformErr := Transform(&message, transform)
	assert.Nil(transformErr)
	assert.Equal(message["a"], "b")
	assert.NotNil(message["x"])
	assert.Equal(message["x"], "y")
}

func TestSimplestNestedTransform(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
	}
	transform := map[string]interface{}{
		"x": "y",
		"c": map[string]interface{}{
			"d": "e",
		},
	}
	transformErr := Transform(&message, transform)
	assert.Nil(transformErr)
	assert.Equal(message["a"], "b")
	assert.NotNil(message["c"])
	assert.NotNil(message["x"])
	assert.Equal(message["x"], "y")
	assert.NotNil(message["c"].(map[string]interface{})["d"])
	assert.Equal(message["c"].(map[string]interface{})["d"], "e")
}
func TestNotSimpleNestedTransform(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"hi": "there",
		},
	}
	transform := map[string]interface{}{
		"x": "y",
		"c": map[string]interface{}{
			"d": "e",
		},
	}
	transformErr := Transform(&message, transform)
	assert.Nil(transformErr)
	assert.Equal(message["a"], "b")
	assert.NotNil(message["c"])
	assert.NotNil(message["x"])
	assert.Equal(message["x"], "y")
	assert.NotNil(message["c"].(map[string]interface{})["hi"])
	assert.Equal(message["c"].(map[string]interface{})["hi"], "there")
	assert.NotNil(message["c"].(map[string]interface{})["d"])
	assert.Equal(message["c"].(map[string]interface{})["d"], "e")
}

func definitelyNo() {
	pretty.Println("no")
}
