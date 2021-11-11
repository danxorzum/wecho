package services

import (
	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

func InitP(projectName string) {
	jio.CreateFolder(viper.GetString("cmddir") + "/" + projectName)
	projectPath := viper.GetString("cmddir") + "/" + projectName
	jio.NewFileforTemplate(projectPath+"/main.go", viper.GetString("homedir")+"/templates/init/main.stub", nil)
}
