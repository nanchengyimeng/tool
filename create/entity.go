package create

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

var text string = `
package entity

type %s struct {

}

type List []*%s


func New%s() *%s {
	return &%s{}
}
`

func CreateEntity(name string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(ProjectDirList[DirDomainEntity])
	path.WriteString("/%s.go")

	fp := fmt.Sprintf(path.String(), name)

	is, err := PathExists(fp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if is {
		fmt.Println("file exist")
		os.Exit(1)
	}

	file, err := os.OpenFile(fp, os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//创建实体文件
	entityName := cases.Title(language.English, cases.NoLower).String(name)
	_, err = file.WriteString(fmt.Sprintf(
		text,
		entityName,
		entityName,
		entityName,
		entityName,
		entityName,
	))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
