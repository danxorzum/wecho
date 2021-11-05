package services

import (
	"fmt"

	"github.com/culturadevops/jgt/jio"
)

type Ctl struct{}

func (t *Ctl) CreateMap(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = fileName
	return MapForReplace
}

func (t *Ctl) GetModelTempleteName(version string) string {
	return "C:/Users/USER/wecho/.config/wecho/models/" + version + ".stub"
}

func (t *Ctl) Mdl(fileName string, version string) {
	FolderBase := "models"
	jio.CreateFolder(FolderBase)
	MapForReplace := t.CreateMap(fileName)
	newName := FolderBase + "/" + fileName + "_model" + ".go"
	fmt.Println(GetModelTempleteName(version))
	jio.NewFileforTemplate(newName, t.GetModelTempleteName(version), MapForReplace)
}
