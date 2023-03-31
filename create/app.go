package create

import (
	"fmt"
	"os"
	"strings"
)

var demoApp = `
package application

import (
	"fmt"
)

type IDemo interface {
	// Save demo
	Save(ctx context.Context, demo string) error
}

var _ IDemo = demoApp{}

type demoApp struct {
	
}

func NewDemoApp() IDemo {
	return &demoApp{
	}
}

func (d demoApp) Save(ctx context.Context, demo string) error {
	fmt.Println(demo)

	//这里可以调用domain & infrastructure封装好的代码，控制业务逻辑的走向

	return nil
}
`

func CreateDempApp(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirApplication])
	path.WriteString("/demoApp.go")

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
		demoApp,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
