package cmd

import (
	"fmt"
	"reflect"

	"errors"
	"github.com/bndr/gopencils"
	"github.com/spf13/viper"
	"github.com/twhiston/clitable"
	"log"
	"net/http"
)

//GetAPI returns a gopencils resource ready configured to make calls on
func GetAPI() *gopencils.Resource {
	domain := viper.GetString("hakuna.domain")
	api := gopencils.Api("https://" + domain + ".hakuna.ch/api/v1")
	api.Headers = make(http.Header, 2)
	api.Headers.Add("X-Auth-Token", viper.GetString("hakuna.token"))
	api.Headers.Add("Content-Type", "application/json; charset=utf-8")
	return api
}

func PrintArrayResponse(resp interface{}) {

	switch reflect.TypeOf(resp).Kind() {
	case reflect.Slice:
		table := clitable.New()
		vp := viper.GetBool("vertical_print")
		if vp {
			printVerticalArray(resp, table)
		} else {
			printNormalArray(resp, table)
		}
		table.Print()
	}
}

func printNormalArray(resp interface{}, table *clitable.Table) {
	s := reflect.ValueOf(resp)
	for i := 0; i < s.Len(); i++ {
		if i == 0 {
			table.AddRow(getStructTags(s.Index(i).Interface())...)
		}
		table.AddRow(getStructVals(s.Index(i).Interface())...)
	}
}

func printVerticalArray(resp interface{}, table *clitable.Table) {

	s := reflect.ValueOf(resp)
	if s.Len() <= 0 {
		return
	}
	tags := getStructTags(s.Index(0).Interface())

	data := make([][]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		data[i] = getStructVals(s.Index(i).Interface())
	}
	for k, v := range tags {
		row := make([]string, 1)
		row[0] = v
		for _, d := range data {
			row = append(row, d[k])
		}
		table.AddRow(row...)
	}
}

//PrintResponse takes some kind of API response and tries to render it as a table
func PrintResponse(resp interface{}) {
	table := clitable.New()
	vp := viper.GetBool("vertical_print")
	if vp {
		printVerticalTable(resp, table)
	} else {
		printNormalTable(resp, table)
	}
	table.Print()
}

func printNormalTable(resp interface{}, table *clitable.Table) {
	//Get api response tags for the header
	table.AddRow(getStructTags(resp)...)
	//Get the data for the table body
	table.AddRow(getStructVals(resp)...)
}

func printVerticalTable(resp interface{}, table *clitable.Table) {
	tags := getStructTags(resp)
	data := getStructVals(resp)
	table.AddRow("Name", "Data")
	for k, v := range tags {
		table.AddRow(v, data[k])
	}
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
			tag := val.Type().Field(i).Tag.Get("json")
			if tag != "" {
				data = append(data, tag)
			}

		}
	}
	return data
}

func secondsToHoursAndMinutes(inSeconds int) string {
	hours := inSeconds / (60 * 60)
	remainder := (inSeconds / 60) % 60
	return fmt.Sprintf("%d:%02d", hours, remainder)
}

func testConfig() error {
	if viper.GetString("hakuna.token") == "" || viper.GetString("hakuna.domain") == "" {
		return errors.New("hakuna.token and hakuna.domain must be present in the configuration file")
	}
	if verbose {
		log.Println("Domain:", viper.GetString("hakuna.domain"))
		log.Println("vertical print:", viper.GetString("vertical_print"))
	}
	return nil
}

//HandleError simply logs the error and hard quits the app
func HandleError(err error) {
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatalln(err)
		}
	}
}

func deferredBodyClose(r *gopencils.Resource) {
	err := r.Raw.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
