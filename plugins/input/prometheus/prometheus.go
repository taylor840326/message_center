package prometheus

import (
	"encoding/json"
	"message_center/message"
	"message_center/utils"
	"time"
)

// labels
type Labels map[string]string

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type AlertItem struct {
	Status      string      `json:"status"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
	StartsAt    string      `json:"startsAt"`
	EndsAt      string      `json:"endsAt"`
	FingerPrint string      `json:"fingerprint"`
}

type PrometheusAlertMessage struct {
	Alerts []AlertItem `json:"alerts"`
}

func (prom *AlertItem) ToString() (string, error) {
	val, err := json.Marshal(prom)
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (prom *AlertItem) ToMessage() (message.Message, error) {
	msg := message.Message{}

	if !utils.IsEmptyString(prom.Annotations.Summary) {
		msg.Summary.Default = prom.Annotations.Summary
	}
	if !utils.IsEmptyString(prom.Annotations.Description) {
		msg.Description.Default = prom.Annotations.Description
	}

	if len(prom.Labels) != 0 {
		msg.Labels = make(map[string]message.I10nField, len(prom.Labels))
		for k, v := range prom.Labels {
			msg.Labels[k] = message.I10nField{
				Default: v,
			}
		}
	}
	msg.Status.Default = prom.Status

	st, err := time.Parse(time.RFC3339Nano, prom.StartsAt)
	if err != nil {
		return msg, err
	}
	msg.StartAt = st

	et, err := time.Parse(time.RFC3339Nano, prom.StartsAt)
	if err != nil {
		return msg, err
	}
	msg.EndAt = et
	return msg, nil
}
