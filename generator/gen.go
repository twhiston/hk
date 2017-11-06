package main

// +build ignore

import (
	"text/template"
	"os"
	"github.com/twhiston/hk/cmd"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type GenDataSet struct {
	Pkg     string   `yaml:"pkg,omitempty"`
	Imports []string `yaml:"import,omitempty"`
	Commands struct {
		Get    []GenGetData  `yaml:"get,omitempty"`
		Post   []GenPostData `yaml:"post,omitempty"`
		Put    []GenPostData `yaml:"put,omitempty"`
		Delete []GenGetData  `yaml:"delete,omitempty"`
	} `yaml:"commands"`
}

type GenGetData struct {
	Id            string            `yaml:"id,omitempty"`
	Parent        string            `yaml:"parent,omitempty"`
	ArrayResponse bool              `yaml:"array,omitempty"`
	Description   string            `yaml:"description,omitempty"`
	Long          string            `yaml:"long,omitempty"`
	ResponseType  string            `yaml:"responseType,omitempty"`
	Path          string            `yaml:"path,omitempty"`
	Params        map[string]string `yaml:"params,omitempty"`
}

type GenPostData struct {
	GenGetData             `yaml:",inline"`
	PayloadType     string `yaml:"payloadType,omitempty"`
	PostDataHandler string `yaml:"postDataHandler,omitempty"`
}

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
	//TODO - pass viper options to generator
	//TODO - make payload rendering optional

	file, err := ioutil.ReadFile("cmdManifest.yml")
	cmd.HandleError(err)

	dataset := newGenDataSet()
	err = yaml.Unmarshal(file, dataset)
	cmd.HandleError(err)

	t := template.New("hk")
	t, err = t.ParseGlob("generator/tmpl/*.tmpl")
	cmd.HandleError(err)

	f, err := os.Create("cmd/commands_generated.go")
	cmd.HandleError(err)
	defer f.Close()

	err = t.ExecuteTemplate(f, "file.tmpl", &dataset)
	cmd.HandleError(err)

}
