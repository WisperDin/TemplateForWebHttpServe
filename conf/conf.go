package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

/**
读取配置文件
*/

type tomlFile struct {
	AppEnv string `toml:"appEnv"`
	DBDriver         string `toml:"DBDriver"`
	DBHost         string `toml:"DBHost"`
	DBPort         string `toml:"DBPort"`
	DBUser         string `toml:"DBUser"`
	DBPassword     string `toml:"DBPassword"`
	DBName         string `toml:"DBName"`

	DBTestName         string `toml:"DBTestName"`

	ServerHost     string `toml:"serverHost"`
	ServerPort     string `toml:"serverPort"`
}

var App *tomlFile

func init() {
	App = new(tomlFile)
}

func Init(filePath string) {
	_, err := toml.DecodeFile(RealFilePath(filePath), App)
	if err != nil {
		panic(err)
	}
	fmt.Println(App)
}
