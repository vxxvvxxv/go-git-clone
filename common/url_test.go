package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type checkUrlFolders struct {
	folders []string
	err     error
}

func TestGetUrlFolders(t *testing.T) {
	for urlOriginal, needThis := range map[string]checkUrlFolders{
		"git@github.com:vxxvvxxv/go-git-clone.git":     {[]string{"github.com", "vxxvvxxv"}, nil},
		"https://github.com/vxxvvxxv/go-git-clone.git": {[]string{"github.com", "vxxvvxxv"}, nil},
	} {
		checkFolders, checkErr, _ := GetUrlFolders(urlOriginal)
		assert.Equal(t, needThis.err, checkErr)
		assert.Equal(t, needThis.folders, checkFolders)
	}
}
