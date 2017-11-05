package config

import (
	"strings"
	"github.com/zouyx/jodz/utils"
)

const appConfigFileName  ="app.properties"

var (
	//appconfig
	appConfig *AppConfig

)

func GetAppConfig() *AppConfig {
	return appConfig
}

type AppConfig struct {
	Zk string `json:"zk"`
	Jobs string `json:"jobs"`
	zkIps []string
	jobNodes []string
}

func (this *AppConfig)GetZkIps() []string {
	if utils.IsNotNil(this.zkIps){
		return this.zkIps
	}

	this.zkIps=this.cutString(this.Zk)

	return this.zkIps
}

func (this *AppConfig)GetJobNodes() []string {
	if utils.IsNotNil(this.jobNodes){
		return this.jobNodes
	}

	this.jobNodes=this.cutString(this.Jobs)

	return this.jobNodes
}

func (this *AppConfig)cutString(str string) []string {
	if utils.IsEmpty(str){
		return []string{}
	}

	return strings.Split(this.Zk,",")
}


func init() {
	//init config
	initConfig()
}

func initConfig() {
	var err error
	//init config file
	appConfig,err = loadJsonConfig(appConfigFileName)

	if err!=nil{
		panic(err)
	}
}