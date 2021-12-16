package public

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

// InitConfig 通常这个获取配置的位置在cobra的root的init函数中
func InitConf() {
	// 使用 go-bindata -o=./public/config.go -pkg=public .env 将配置文件转成go文件,然后使用如下方式将配置读取,然后将.env加入到 .gitignore中,即可实现密码不上传到git仓库,部署到服务器亦可正常使用
	// 当前为了方便使用,直接使用了配置文件,如果需要使用配置文件,请使用上面的方式
	fileobj, err := ioutil.ReadFile(".env")
	if err != nil {
		fmt.Printf("Asset file err: %v\n", err)
		return
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(fileobj))
	viper.AutomaticEnv() // read in environment variables that match
}

func GetRunEvn() (env string) {
	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "test" {
		env = "test"
	}
	if os.Getenv("ENV") == "prod" {
		env = "prod"
	}
	return
}
