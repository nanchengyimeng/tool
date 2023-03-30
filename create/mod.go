package create

import (
	"fmt"
	"os"
	"strings"
)

var modText = `
module %s

go 1.20
`

func createMod(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/go.mod")

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
		modText,
		projectName,
	))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
