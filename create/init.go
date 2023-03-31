package create

import (
	"fmt"
	"strings"
)

const (
	DirApplication = iota
	DirCmd
	DirConf
	DirDomain
	DirDomainAgg
	DirDomainEntity
	DirDomainModel
	DirDomainResp
	DirDomainRespCache
	DirDomainRespLib
	DirDomainRespPer
	DirDomainSrv

	DirInf
	DirInfCache
	DirInfLib
	DirInfPer

	DirInterface
	DirInterfaceMiddle
	DirInterfaceCtl
	DirInterfaceRpc

	DirRoute
)

var ProjectDirList = map[int]string{
	DirApplication: "application",
	DirCmd:         "cmd",
	DirConf:        "config",

	DirDomain:          "domain",
	DirDomainAgg:       "domain/aggregate",
	DirDomainEntity:    "domain/entity",
	DirDomainModel:     "domain/model",
	DirDomainResp:      "domain/repository",
	DirDomainRespCache: "domain/repository/cache",
	DirDomainRespLib:   "domain/repository/library",
	DirDomainRespPer:   "domain/repository/persistence",
	DirDomainSrv:       "domain/services",

	DirInf:      "infrastructure",
	DirInfCache: "infrastructure/cache",
	DirInfLib:   "infrastructure/library",
	DirInfPer:   "infrastructure/persistence",

	DirInterface:       "interfaces",
	DirInterfaceMiddle: "interfaces/middleware",
	DirInterfaceCtl:    "interfaces/controller",
	DirInterfaceRpc:    "interfaces/rpc",

	DirRoute: "router",
}

type project struct {
	Name string
}

func newProject(name string) *project {
	return &project{Name: name}
}

func InitProject(name string) {
	p := newProject(name)
	err := CreateMultiDir("./" + name)
	if err != nil {
		panic(err)
	}

	for _, dir := range ProjectDirList {
		if err := p.init(dir); err != nil {
			panic(err)
		}
	}

	//创建路由
	CreateRouter(name)
	//创建示例ctl
	CreateDemoController(name)
	//创建示例应用
	CreateDempApp(name)

	CreateInf(name)
	createInfCache(name)
	createInfLib(name)
	createInfPer(name)
	createInfPerBase(name)

	createCmd(name)

	createMod(name)

	echo := `
project create success
please run 'go mod tidy' in your project!`
	fmt.Println(echo)
}

func (p *project) init(mode string) error {
	var path strings.Builder

	path.WriteString("./")
	path.WriteString(p.Name)
	path.WriteByte('/')
	path.WriteString(mode)

	err := CreateMultiDir(path.String())
	if err != nil {
		return err
	}
	return nil
}
