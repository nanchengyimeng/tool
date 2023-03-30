package create

import (
	"fmt"
	"os"
	"strings"
)

var cmdTxt = `
package main

import (
	"%s/router"

	_ "git.ruigushop.com/golang/rgo"
	"git.ruigushop.com/golang/rgo/core/rgrouter"
	_ "git.ruigushop.com/golang/rgo/util/rgstarthook"
)

func main() {
	rgrouter.Run(router.GetRouter())
}
`

func createCmd(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirCmd])
	path.WriteString("/main.go")

	is, err := PathExists(path.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if is {
		fmt.Println("file exist")
		os.Exit(1)
	}

	file, err := os.OpenFile(path.String(), os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//创建实体文件
	_, err = file.WriteString(fmt.Sprintf(
		cmdTxt,
		projectName,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
