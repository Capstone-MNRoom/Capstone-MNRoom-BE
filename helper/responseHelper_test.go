package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseHelpers(t *testing.T) {
	t.Run("ResponseFailed", func(t *testing.T) {
		res := ResponseFailed("bad request")
		assert.Equal(t, "error", res["status"])
		assert.Equal(t, "bad request", res["message"])
	})

	t.Run("ResponseSuccessNoData", func(t *testing.T) {
		res := ResponseSuccessNoData("operation succeeded")
		assert.Equal(t, "success", res["status"])
		assert.Equal(t, "operation succeeded", res["message"])
	})

	t.Run("ResponseSuccessWithData", func(t *testing.T) {
		data := map[string]string{"key": "value"}
		res := ResponseSuccessWithData("data found", data)
		assert.Equal(t, "success", res["status"])
		assert.Equal(t, "data found", res["message"])
		assert.Equal(t, data, res["data"])
	})
}
