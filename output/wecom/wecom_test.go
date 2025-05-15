package wecom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"message_center/input/prometheus"
	"os"
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {

	c, err := ioutil.ReadFile("./body.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	message := prometheus.AlertMessage{}
	err = json.Unmarshal(c, &message)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	alerts := message.Alerts
	rval := reflect.ValueOf(alerts[0])

	fmt.Println(rval.NumField())
	fmt.Println(rval.Field(0))
	fmt.Println(rval.FieldByName("Labels"))
}
