package config

import (
	"encoding/json"
	"log"
	"mygame/config/enum"
	"mygame/package/rootpath"

	"github.com/spf13/viper"
)

const (
	default_ProjectName = "mygame"
	default_ServiceName = "mygame"
	default_Environment = enum.Environment_Local
)

var (
	Config = config{Viper: viper.New()}
)

type config struct {
	*viper.Viper
	System system
	Server server
}

type system struct {
	ServiceName     string
	Environment     string
	ProjectRootPath string
}

type server struct {
	Http struct{ Port string }
	GRPC struct{ Port string }
}

func init() {
	// Environment
	Config.AutomaticEnv()
	Config.SetDefault("SERVICE", default_ServiceName)
	Config.SetDefault("ENV", default_Environment)
	Config.System.ServiceName = Config.Viper.GetString("SERVICE")
	Config.System.Environment = Config.Viper.GetString("ENV")
	// Root Path
	path, err := rootpath.GetFilePath(default_ProjectName)
	if err != nil {
		log.Panic("[取得專案根目錄錯誤]: ", err)
	}
	Config.System.ProjectRootPath = path
	// Yaml
	Config.SetConfigName(Config.System.Environment)
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./yaml")
	Config.AddConfigPath("./config/yaml")
	Config.AddConfigPath(Config.System.ProjectRootPath + "/config/yaml")
	// Read
	if err := Config.ReadInConfig(); err != nil {
		log.Panic("[設定檔讀取錯誤]: ", err)
	}
	// Unmarshal
	if err := Config.Unmarshal(&Config); err != nil {
		log.Panic("[設定檔解析錯誤]: ", err)
	}
}

func (c *config) Print() string {
	b, _ := json.MarshalIndent(c, "", "    ")
	return string(b)
}
