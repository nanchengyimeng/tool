package create

import (
	"fmt"
	"os"
)

func Model(name string) {
	path := "./domain/entity/%s.go"
	path = fmt.Sprintf(path, name)

	os.OpenFile(path, os.O_CREATE, os.ModePerm)
	//创建实体文件
}
