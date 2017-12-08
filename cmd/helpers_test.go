package cmd

import (
	"bytes"
	"errors"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestHandleErrorNil(t *testing.T) {
	HandleError(nil)
}

func TestHandleError(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		HandleError(errors.New("test error"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestHandleError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}

func TestHandleEofError(t *testing.T) {
	inErr := errors.New("EOF")
	HandleError(inErr)
}

func TestTestConfig(t *testing.T) {
	err := testConfig()
	if err == nil {
		t.Fatal("Invalid config should return an error")
	}

	viper.Set("hakuna.token", "test")

	err = testConfig()
	if err == nil {
		t.Fatal("Invalid config should return an error")
	}

	viper.Set("hakuna.domain", "test")
	err = testConfig()
	if err != nil {
		t.Fatal("Valid config should not return an error")
	}

	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	verbose = true
	err = testConfig()
	if err != nil {
		t.Fatal("Valid config should not return an error")
	}
	if buf.String() != "Domain: test\n" {
		t.Fatal("incorrect log output", buf.String())
	}

}

func TestGetApi(t *testing.T) {
	// Not a lot to test with this as it just depends on having the viper credentials or will die at runtime
	api := GetAPI()
	if reflect.TypeOf(api).String() != "*gopencils.Resource" {
		t.Fatal("API must be a gopencils resource")
	}
}

type TestStruct struct {
	Name  string `json:"name"`
	Value int
	Nest  struct {
		Name  string
		Value int `json:"val"`
	}
}

func TestGetStructVals(t *testing.T) {
	rand.Seed(time.Now().Unix())
	ts := new(TestStruct)
	for i := 0; i < 666; i++ {
		ts.Name = RandStringBytesMaskImprSrc(rand.Int() % 32)
		ts.Value = rand.Int()
		ts.Nest.Name = RandStringBytesMaskImprSrc(rand.Int() % 32)
		ts.Nest.Value = rand.Int()
		res := getStructVals(*ts)
		if res[0] != ts.Name || res[1] != strconv.Itoa(ts.Value) || res[2] != ts.Nest.Name || res[3] != strconv.Itoa(ts.Nest.Value) {
			t.Fatal("struct values not correctly determined", ts)
		}
	}
}

func TestGetStructTags(t *testing.T) {

	ts := new(TestStruct)
	tags := getStructTags(*ts)
	if tags[0] != "name" || tags[1] != "val" {
		t.Fatal("struct tags were not correctly retrieved", "struct", ts, "tags", tags)
	}
}

type TimePair struct {
	Seconds     int
	Stringified string
}

func TestSecondsToHoursAndMinutes(t *testing.T) {

	times := []TimePair{
		{60, "0:01"},
		{600, "0:10"},
		{86400, "24:00"},
		{5400, "1:30"},
	}

	for _, v := range times {
		out := secondsToHoursAndMinutes(v.Seconds)
		if out != v.Stringified {
			t.Fatal("time was not calculated correctly", out, v.Stringified)
		}
	}

}

// HELPER FUNCTIONS

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
