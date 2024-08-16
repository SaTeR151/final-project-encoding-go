package encoding

import (
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"

	"encoding/json"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	// ...
	jFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jFile, &j.DockerCompose)
	if err != nil {
		return err
	}
	yFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	yData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}
	_, err = yFile.Write(yData)
	if err != nil {
		return err
	}
	defer yFile.Close()
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, &y.DockerCompose)
	if err != nil {
		return err
	}
	yamlFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return err
	}
	return nil
}
