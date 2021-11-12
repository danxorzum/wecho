package services

import (
	"log"
	"os/exec"

	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

func InitP(projectName string) {
	jio.CreateFolder(viper.GetString("cmddir") + "/" + projectName)
	projectPath := viper.GetString("cmddir") + "/" + projectName
	jio.NewFileforTemplate(projectPath+"/main.go", viper.GetString("homedir")+"/templates/init/main.stub", nil)

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = viper.GetString("cmddir") + "/" + projectName
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	tidy := exec.Command("go", "mod", "tidy")
	tidy.Dir = viper.GetString("cmddir") + "/" + projectName
	err = tidy.Run()
	if err != nil {
		log.Fatal(err)
	}
}
