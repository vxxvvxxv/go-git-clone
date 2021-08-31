package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vxxvvxxv/go-git-clone/common"
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

	err = os.Chdir(pathProject)
	common.CheckIfError(err, common.ErrCreateFolders)

	cmd := exec.Command("git", "clone", originUrl)
	err = cmd.Run()
	common.CheckIfError(err, common.ErrClone)

	common.Info("Success cloned: %s", originUrl)
}
