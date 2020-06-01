package photos

import (
	"memoriesbot/pkg/config"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	c := config.Get()
	res, err := MakeRequest()

	t.Log(c.Auth.GooglePhotos.AccessToken)

	if err != nil {
		t.Error(err)
	}

	t.Error(res)
}
