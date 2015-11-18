package messagetransform

import (
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
