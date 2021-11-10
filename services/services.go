package services

import (
	"strings"

	"github.com/culturadevops/jgt/jio"
	"github.com/spf13/viper"
)

var Srvs *Services

type Services struct{}

func (t *Services) aUpper(b string) string {
	b = strings.ToLower(b)
	b = strings.Title(b)
	return b
}

func (t *Services) GetTemplateName(folderName string, filesName string) string {
	return viper.GetString("homedir") + "/templates/" + folderName + "/" + filesName + ".stub"
}
func (t *Services) CreateMapDefault(fileName string) map[string]string {
	MapForReplace := make(map[string]string)

	MapForReplace["%MODELNAME%"] = t.aUpper(fileName)
	MapForReplace["%MODELLOWNAME%"] = fileName
	return MapForReplace
}

func (t *Services) TestCreateFile(folderName string, filesName string, templateName string) {
	tempName := t.GetTemplateName(folderName, templateName)
	currentP := viper.GetString("cmddir") + "/"

	jio.CreateFolder(currentP + folderName)
	MapForReplace := t.CreateMapDefault(filesName)
	newName := currentP + folderName + "/" + filesName + "_" + folderName + ".go"

	// file, _ := ioutil.ReadFile(tempName)
	// fmt.Println(file)

	jio.NewFileforTemplate(newName, tempName, MapForReplace)

}

// func (t *Service) CreateDefaultFile(origFolderName string, filesVersionName string, destFolderName string, nameStruct string, destFileName string) {
// 	path, _ := filepath.Abs("")
// 	jio.CreateFolder(path + "/lib/" + origFolderName)
// 	MapForReplace := t.CreateMapDefault(nameStruct)
// 	newName := path + "/lib/" + destFolderName + "/" + destFileName
// 	//fmt.Println(newName)
// 	fmt.Println("C:/Users/USER/")
// 	//fmt.Println(MapForReplace)
// 	tes, _ := ioutil.ReadFile(viper.GetString("homedir") + "./templates/" + origFolderName + "/" + filesVersionName + ".stub")
// 	fmt.Println(tes)

// 	jio.NewFileforTemplate(newName, t.GetTempleteName(origFolderName, filesVersionName), MapForReplace)
// }
