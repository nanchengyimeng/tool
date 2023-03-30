package create

import (
	"fmt"
	"os"
	"strings"
)

var routerText = `
package router

import (
	"%s/application"
	"%s/infrastructure"
	"%s/interfaces/controller"

	"git.ruigushop.com/golang/rgo/core/rgrouter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	router := rgrouter.NewRouter()

	// 获取基础层
	repo := infrastructure.NewRepo()

	_ = repo

	/************************ Demo 示例代码***********************/
	//获取应用
	demoApp := application.NewDemoApp()
	//应用注入控制器
	demoController := controller.NewDemoHandler(demoApp)
	//编写路由
	demoGroup := router.Group("/demo")
	{
		demoGroup.GET("/demo", demoController.Demo)
	}

	/************************ 业务代码    *************************/

	/************************************************************/
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "我见青山多妩媚，料青山见我应如是 --- qiu \n")
		return
	})
	return router
}

`

func CreateRouter(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirRoute])
	path.WriteString("/router.go")

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
		routerText,
		projectName,
		projectName,
		projectName,
	))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
