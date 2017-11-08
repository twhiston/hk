//+build ignore

package main

import (
	"github.com/twhiston/hk/cmd"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type genDataSet struct {
	Pkg      string   `yaml:"pkg,omitempty"`
	Imports  []string `yaml:"import,omitempty"`
	Commands struct {
		Get    []genGetData  `yaml:"get,omitempty"`
		Post   []genPostData `yaml:"post,omitempty"`
		Put    []genPostData `yaml:"put,omitempty"`
		Delete []genGetData  `yaml:"delete,omitempty"`
	} `yaml:"commands"`
}

type genGetData struct {
	ID               string                `yaml:"id,omitempty"`
	Parent           string                `yaml:"parent,omitempty"`
	ArrayResponse    bool                  `yaml:"array,omitempty"`
	Description      string                `yaml:"description,omitempty"`
	Long             string                `yaml:"long,omitempty"`
	ResponseType     string                `yaml:"responseType,omitempty"`
	Path             string                `yaml:"path,omitempty"`
	ParameterHandler string                `yaml:"parameterHandler,omitempty"`
	Params           []genCommandParameter `yaml:"params,omitempty"`
	HasResponse      bool                  `yaml:"hasResponse,omitempty"`
	Index            genIndexData          `yaml:"index,omitempty"`
	ExpectedStatus   int                   `yaml:"expectedStatus"`
}

type genIndexData struct {
	Arg      int  `yaml:"index,omitempty"`
	Required bool `yaml:"required,omitempty"`
}

type genPostData struct {
	genGetData      `yaml:",inline"`
	PayloadType     string `yaml:"payloadType,omitempty"`
	PostDataHandler string `yaml:"postDataHandler,omitempty"`
}

type genCommandOption struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Usage      string `yaml:"usage"`
	Value      string `yaml:"value"`
	Persistent bool   `yaml:"persistent"`
	Short      string `yaml:"short,omitempty"`
}

type genCommandParameter genCommandOption

func newGenDataSet() *genDataSet {
	gds := new(genDataSet)
	gds.Pkg = "cmd" //TODO - make dynamic
	gds.Imports = make([]string, 0)
	gds.Commands.Delete = make([]genGetData, 0)
	gds.Commands.Put = make([]genPostData, 0)
	gds.Commands.Post = make([]genPostData, 0)
	gds.Commands.Get = make([]genGetData, 0)
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
	defer func(r *os.File) {
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	err = t.ExecuteTemplate(f, "file.tmpl", &dataset)
	cmd.HandleError(err)

	formatter := exec.Command("gofmt", "-s", "-w", "cmd/commands_generated.go")
	cmd.HandleError(formatter.Run())

}
