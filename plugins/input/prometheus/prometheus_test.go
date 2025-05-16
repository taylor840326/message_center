package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAlertItemToString(t *testing.T) {
	content, err := ioutil.ReadFile("./body.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prometheus_message := PrometheusAlertMessage{}
	err = json.Unmarshal(content, &prometheus_message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	msg, err := prometheus_message.ToMessages()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(msg)
}
