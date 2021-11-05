package services

import (
	"fmt"

	"github.com/culturadevops/jgt/jio"
)

func CreateMap(fileName string) map[string]string {
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = fileName
	return MapForReplace
}

func GetModelTempleteName(version string) string {
	return "C:/Users/USER/wecho/.config/wecho/models/" + version + ".stub"
}

func Mdl(fileName string, version string) {
	FolderBase := "models"
	jio.CreateFolder(FolderBase)
	MapForReplace := CreateMap(fileName)
	newName := FolderBase + "/" + fileName + "_model" + ".go"
	fmt.Println(GetModelTempleteName(version))
	jio.NewFileforTemplate(newName, GetModelTempleteName(version), MapForReplace)
}
