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

func (t *Services) getTemplateName(folderName string, filesName string) string {
	return viper.GetString("homedir") + "/templates/" + folderName + "/" + filesName + ".stub"
}
func (t *Services) createMapDefault(fileName string) map[string]string {
	MapForReplace := make(map[string]string)

	MapForReplace["%EXPORTNAME%"] = t.aUpper(fileName)
	MapForReplace["%MODELNAME%"] = fileName
	return MapForReplace
}

func (t *Services) CreateFile(folderName string, filesName string, templateName string) {
	tempName := t.getTemplateName(folderName, templateName)
	currentP := viper.GetString("cmddir") + "/"

	jio.CreateFolder(currentP + folderName)
	MapForReplace := t.createMapDefault(filesName)
	newName := currentP + folderName + "/" + filesName + "_" + folderName + ".go"

	// file, _ := ioutil.ReadFile(tempName)
	// fmt.Println(file)

	jio.NewFileforTemplate(newName, tempName, MapForReplace)

}
