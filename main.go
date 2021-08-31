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
	project, folders, err, errMessage := common.GetUrlFolders(originUrl)
	common.CheckIfError(err, errMessage)

	goPathFolders := strings.Split(os.Getenv("GOPATH"), ":")[0]
	listFolders := append([]string{}, goPathFolders, "src")
	listFolders = append(listFolders, folders...)

	common.Info("URL: %s\nPath: %s", originUrl, strings.Join(listFolders, string(os.PathSeparator)))

	var isExistsFolders bool
	var folderToCreate string

	for i := 1; i <= len(listFolders); i++ {
		folderToCreate = strings.Join(listFolders[0:i], string(os.PathSeparator))

		isExistsFolders, err = common.IsExists(folderToCreate)
		common.CheckIfError(err, common.ErrCreateFolders)

		if !isExistsFolders {
			err = os.Mkdir(folderToCreate, 0755)
			common.CheckIfError(err, common.ErrCreateFolders)
		}

		err = os.Chdir(folderToCreate)
		common.CheckIfError(err, common.ErrCreateFolders)
	}

	// Check if exists project

	projectPath := strings.Join(listFolders, string(os.PathSeparator)) + string(os.PathSeparator) + project

	isExistsFolders, err = common.IsExists(projectPath)
	common.CheckIfError(err, common.ErrCreateFolders)

	if isExistsFolders {
		common.Info("Already up to date: %s", projectPath)
		os.Exit(0)
	}

	cmd := exec.Command("git", "clone", originUrl)
	err = cmd.Run()
	common.CheckIfError(err, common.ErrClone)

	common.Info("Success cloned: %s", originUrl)
}
