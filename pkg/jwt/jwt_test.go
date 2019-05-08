package jwt

import (
	"testing"
)

func TestGenerate(t *testing.T) {

	t.Run("Generate JWT", func(t *testing.T) {
		_, err := Generate("eu")
		if err != nil {
			t.Error(err)
		}
	})
}
