package args

import (
	"flag"
	"fmt"
	"os"
	"tool/create"
)

type Arguments struct {
	create *string
	entity *string
}

func (a *Arguments) parseArgs(args []string) error {
	a.create = flag.String("create", "", "-create projectName")
	a.entity = flag.String("entity", "", "-entity projectName")

	flag.Parse()
	return nil
}

func (a *Arguments) Execute() error {
	if err := a.parseArgs(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//调用对应处理
	if len(*a.create) != 0 {
		create.InitProject(*a.create)
	}

	if len(*a.entity) != 0 {
		create.CreateEntity(*a.entity)
	}
	return nil
}
