package common

import (
	"errors"
	"net/url"
	"strings"
)

func GetUrlFolders(originUrl string) (string, []string, error, string) {
	if strings.HasPrefix(originUrl, "git@") {
		return getFoldersFromSSH(originUrl)
	}
	if strings.HasPrefix(originUrl, "http://") || strings.HasPrefix(originUrl, "https://") {
		return getFoldersFromHttps(originUrl)
	}

	return "", []string{}, errors.New(originUrl), ErrURL
}

func getFoldersFromSSH(urlString string) (string, []string, error, string) {
	fromIdx := strings.Index(urlString, "@")
	toIdx := strings.Index(urlString, ".git")

	// git@github.com:vxxvvxxv/go-git-clone.git -> git://github.com/vxxvvxxv/go-git-clone
	return getFoldersFromHttps("git://" + strings.ReplaceAll(urlString, ":", "/")[fromIdx+1:toIdx])
}

func getFoldersFromHttps(urlString string) (string, []string, error, string) {
	urlParse, err := url.Parse(urlString)
	if err != nil {
		return "", []string{}, err, ErrURL
	}

	folders := []string{urlParse.Host}
	pathFolders := strings.Split(urlParse.Path, "/")

	if len(pathFolders) <= 1 {
		return "", folders, errors.New("folders is missing"), ErrURL
	}

	folders = append(folders, pathFolders[1:len(pathFolders)-1]...)
	project := strings.Split(pathFolders[len(pathFolders)-1:][0], ".")[0]

	return project, folders, nil, ""
}
