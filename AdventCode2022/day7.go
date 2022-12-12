package main

import (
	"log"
	"strconv"
	"strings"
)

type Folder struct {
	subfolders   []string
	parentfolder string
	totsize      int
}

func Puzzle7_1(filename string) int {

	commands := ReadFiletoStrArray(filename)
	folders := make(map[string]Folder, len(commands))
	var currfolderstr string
	for i := 0; i < len(commands); i++ {
		line := commands[i]
		// this occurs only once.
		arr := strings.Split(line, " ")
		if strings.HasPrefix(line, "$ cd /") {
			currfolderstr = "root"
			folders[currfolderstr] = Folder{parentfolder: "", totsize: 0, subfolders: make([]string, 0, len(commands))}
		} else if strings.HasPrefix(line, "$ cd ..") {
			currfolderstr = folders[currfolderstr].parentfolder
			log.Println("Current folder :", currfolderstr)
		} else if strings.HasPrefix(line, "$ cd") {
			parentfold := currfolderstr
			currfolderstr = arr[2] // new directory, so new current folder
			_, isPresent := folders[currfolderstr]
			if !isPresent {
				folder := Folder{parentfolder: parentfold, totsize: 0, subfolders: make([]string, 0, len(commands))}
				folders[currfolderstr] = folder
			}
		}
		if line == "$ ls" {
			continue
		}
		// check if line is dir
		if strings.HasPrefix(line, "dir") {
			currf, _ := folders[currfolderstr]
			currf.subfolders = append(currf.subfolders, arr[1])
			folders[currfolderstr] = currf
			// add the new directory to map
			_, isPresent := folders[arr[1]]
			if !isPresent {
				nfolder := Folder{parentfolder: currfolderstr, totsize: 0, subfolders: make([]string, 0, len(commands))}
				folders[arr[1]] = nfolder
			}
		}
		sz, err := strconv.Atoi(arr[0])
		if err == nil {
			updateAllParentFolderSize(folders, sz, currfolderstr)
		}

	}
	//fmt.Println(folders)
	sum := 0
	for _, folder := range folders {
		foldersz := folder.totsize
		if foldersz < 100000 {
			sum += foldersz
		}
	}
	return sum
}

func updateAllParentFolderSize(folders map[string]Folder, sz int, currfolderstr string) {
	currf, _ := folders[currfolderstr]
	currf.totsize += sz
	folders[currfolderstr] = currf

	for {
		if currf.parentfolder != "" {
			currparent := currf.parentfolder
			currf = folders[currparent]
			currf.totsize += sz
			folders[currparent] = currf
		} else {
			break
		}
	}
}
