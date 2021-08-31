package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vxxvvxxv/go-git-clone/common"

	"github.com/go-git/go-git/v5"
)

func main() {
	args := os.Args

	if len(args) != 2 || strings.TrimSpace(args[1]) == "" {
		common.CheckIfError(fmt.Errorf("%v", args[1:]), common.ErrArgs)
	}
	originUrl := strings.TrimSpace(args[1])
	folders, err, errMessage := common.GetUrlFolders(originUrl)
	common.CheckIfError(err, errMessage)

	goPathFolders := strings.Split(os.Getenv("GOPATH"), ":")[0]
	pathFolders := append([]string{}, goPathFolders, "src")
	pathFolders = append(pathFolders, folders...)
	pathProject := strings.Join(pathFolders, string(os.PathSeparator))

	common.Info("URL: %s\nPath: %s", originUrl, pathFolders)

	var isExistsFolders bool

	isExistsFolders, err = common.IsExists(pathProject)
	common.CheckIfError(err, common.ErrCreateFolders)

	if !isExistsFolders {
		err = os.Mkdir(pathProject, 0755)
		common.CheckIfError(err, common.ErrCreateFolders)
	}

	_, err = git.PlainClone(pathProject, false, &git.CloneOptions{
		URL:      originUrl,
		Progress: os.Stdout,
	})
	common.CheckIfError(err, common.ErrClone)

	common.Info("Success cloned: %s", originUrl)
}
