package cmd

import (
	"reflect"
	"fmt"

	"github.com/twhiston/clitable"
	"github.com/spf13/viper"
	"github.com/bndr/gopencils"
	"net/http"
	"log"
)

func GetApi() *gopencils.Resource {
	domain := viper.GetString("hakuna.domain")
	api := gopencils.Api("https://" + domain + ".hakuna.ch/api/v1/")
	api.Headers = make(http.Header, 2)
	api.Headers.Add("X-Auth-Token", viper.GetString("hakuna.token"))
	api.Headers.Add("Content-Type", "application/json; charset=utf-8")
	return api
}

func PrintResponse(resp interface{}) {
	table := clitable.New()
	//Get api response tags for the header
	table.AddRow(getStructTags(resp)...)
	//Get the data for the table body
	table.AddRow(getStructVals(resp)...)
	table.Print()
}

func getStructVals(resp interface{}) []string {
	data := make([]string, 0)
	val := reflect.ValueOf(resp)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Type.Kind() == reflect.Struct {
			a := getStructVals(val.Field(i).Interface())
			data = append(data, a...)
		} else {
			data = append(data, fmt.Sprintf("%v", val.Field(i).Interface()))
		}
	}
	return data
}

func getStructTags(resp interface{}) []string {
	data := make([]string, 0)
	val := reflect.ValueOf(resp)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Type.Kind() == reflect.Struct {
			a := getStructTags(val.Field(i).Interface())
			data = append(data, a...)
		} else {
			data = append(data, val.Type().Field(i).Tag.Get("json"))
		}
	}
	return data
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
