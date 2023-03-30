package create

import (
	"fmt"
	"os"
	"strings"
)

var infText = `
package infrastructure

import (
	"%s/infrastructure/cache"
	"%s/infrastructure/library"
	"%s/infrastructure/persistence"
)

type Repo struct {
	Cache       *cache.Repositories
	Persistence *persistence.Repositories
	Library     *library.Library
}

func NewRepo() *Repo {
	return &Repo{
		Cache:       cache.NewRepositories(),
		Persistence: persistence.NewPersistence(),
		Library:     library.NewLibrary(),
	}
}

`

func CreateInf(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirInf])
	path.WriteString("/infrastructure.go")

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
		infText,
		projectName,
		projectName,
		projectName,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var infCacheText = `
package cache

import (
	"git.ruigushop.com/golang/rgo/core/rgmodel/rgredis"
	"github.com/go-redis/redis"
)

type Repositories struct {
	redis redis.Cmdable
}

func NewRepositories() *Repositories {
	client, err := rgredis.GetClient()
	if err != nil {
		panic("初始化缓存仓库失败：" + err.Error())
	}

	return &Repositories{
		redis: client,
	}
}

`

func createInfCache(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirInfCache])
	path.WriteString("/base.go")

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
		infCacheText,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var infLibTxt = `
package library

import (
	"github.com/nanchengyimeng/ghttp"
)

type Library struct {
	//TODO 加入各个三方请求实体

	build *ghttp.ClientBuilder
}

func NewLibrary() *Library {
	build := ghttp.NewClientBuilder()


	return &Library{
		MpServer: mp.NewServer(build),
		Cc:       cc.NewServer(build),
		Order:    order.NewServer(build),
		build:    build,
	}
}

`

func createInfLib(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirInfLib])
	path.WriteString("/base.go")

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
		infLibTxt,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var infPerTxt = `
package persistence

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Db *gorm.DB //对外暴露Db，允许自行使用
	tx *gorm.DB //对外不暴露事务db，防止瞎搞

}

func newRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Db: db,
		tx: db,
	}
}

func NewPersistence() *Repositories {
	var db *gorm.DB
	return newRepositories(db)
}

func (r *Repositories) Begin() *Repositories {
	return newRepositories(r.Db.Begin())
}

func (r *Repositories) Rollback() *gorm.DB {
	return r.tx.Rollback()
}

func (r *Repositories) Commit() *gorm.DB {
	return r.tx.Commit()
}
`

func createInfPer(projectName string) {
	var path strings.Builder
	path.WriteString("./")
	path.WriteString(projectName)
	path.WriteString("/")
	path.WriteString(ProjectDirList[DirInfPer])
	path.WriteString("/base.go")

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
		infPerTxt,
	))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
