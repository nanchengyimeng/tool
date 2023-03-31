package create

import (
	"fmt"
	"os"
	"strings"
)

var controllerText = `
package controller

import (
	"%s/application"
	"github.com/gin-gonic/gin"
)

type demoHandler struct {
	demoApp application.IDemo
}

func NewDemoHandler(d application.IDemo) *demoHandler {
	return &demoHandler{
		demoApp: d,
	}
}

func (d *demoHandler) Demo(c *gin.Context) {
	//在控制器里，不许有业务逻辑代码，只能做数据验证，要完全保证应用层传入的数据是对的

	d.demoApp.Save(c, "秋")
}
`

func CreateDemoController(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirInterfaceCtl])
	path.WriteString("/demoHandle.go")

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
		controllerText,
		projectName,
	))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
