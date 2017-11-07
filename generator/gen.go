//+build ignore

package main

import (
	"github.com/twhiston/hk/cmd"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type GenDataSet struct {
	Pkg      string   `yaml:"pkg,omitempty"`
	Imports  []string `yaml:"import,omitempty"`
	Commands struct {
		Get    []GenGetData  `yaml:"get,omitempty"`
		Post   []GenPostData `yaml:"post,omitempty"`
		Put    []GenPostData `yaml:"put,omitempty"`
		Delete []GenGetData  `yaml:"delete,omitempty"`
	} `yaml:"commands"`
}

type GenGetData struct {
	Id               string                `yaml:"id,omitempty"`
	Parent           string                `yaml:"parent,omitempty"`
	ArrayResponse    bool                  `yaml:"array,omitempty"`
	Description      string                `yaml:"description,omitempty"`
	Long             string                `yaml:"long,omitempty"`
	ResponseType     string                `yaml:"responseType,omitempty"`
	Path             string                `yaml:"path,omitempty"`
	ParameterHandler string                `yaml:"parameterHandler,omitempty"`
	Params           []GenCommandParameter `yaml:"params,omitempty"`
	HasResponse      bool                  `yaml:"hasResponse,omitempty"`
	Index            GenIndexData          `yaml:"index,omitempty"`
	ExpectedStatus   int                   `yaml:"expectedStatus"`
}

type GenIndexData struct {
	Arg      int  `yaml:"index,omitempty"`
	Required bool `yaml:"required,omitempty"`
}

type GenPostData struct {
	GenGetData      `yaml:",inline"`
	PayloadType     string `yaml:"payloadType,omitempty"`
	PostDataHandler string `yaml:"postDataHandler,omitempty"`
}

type GenCommandOption struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Usage      string `yaml:"usage"`
	Value      string `yaml:"value"`
	Persistent bool   `yaml:"persistent"`
	Short      string `yaml:"short,omitempty"`
}

type GenCommandParameter GenCommandOption

func newGenDataSet() *GenDataSet {
	gds := new(GenDataSet)
	gds.Pkg = "cmd" //TODO - make dynamic
	gds.Imports = make([]string, 0)
	gds.Commands.Delete = make([]GenGetData, 0)
	gds.Commands.Put = make([]GenPostData, 0)
	gds.Commands.Post = make([]GenPostData, 0)
	gds.Commands.Get = make([]GenGetData, 0)
	return gds
}

func main() {

	//TODO - put template is basically post. could refactor this to use the same one?
	//TODO - optional dynamic generation of payload renderer stubs

	file, err := ioutil.ReadFile("manifest.yml")
	cmd.HandleError(err)

	dataset := newGenDataSet()
	err = yaml.Unmarshal(file, dataset)
	cmd.HandleError(err)

	t := template.New("hk")

	funcMap := template.FuncMap{
		"ToUpper": strings.Title,
	}
	t.Funcs(funcMap)

	t, err = t.ParseGlob("generator/tmpl/*.tmpl")
	cmd.HandleError(err)

	f, err := os.Create("cmd/commands_generated.go")
	cmd.HandleError(err)
	defer f.Close()

	err = t.ExecuteTemplate(f, "file.tmpl", &dataset)
	cmd.HandleError(err)

}
