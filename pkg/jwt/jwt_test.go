package jwt

import (
	"testing"
)
func TestGenerateJWT(t *testing.T){
	token, err := GenerateJWT("eu")
	if err != nil {
		t.Error(err)
	}
	t.Error(token)
}