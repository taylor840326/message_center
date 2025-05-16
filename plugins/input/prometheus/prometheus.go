package prometheus

import (
	"encoding/json"
	"message_center/message"
	"message_center/utils"
	"strings"
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

func NewPrometheusInputProcessor(context []byte) (*PrometheusAlertMessage, error) {
	msg := &PrometheusAlertMessage{}
	err := json.Unmarshal(context, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (prom *PrometheusAlertMessage) ToString() (string, error) {
	val, err := json.Marshal(prom)
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (prom *PrometheusAlertMessage) ToMessages() ([]message.Message, error) {
	alerts := prom.Alerts
	messages := make([]message.Message, 0)

	for _, alert := range alerts {
		msg := message.Message{}
		if !utils.IsEmptyString(alert.Annotations.Summary) {
			msg.Summary.Default = alert.Annotations.Summary
		}
		if !utils.IsEmptyString(alert.Annotations.Description) {
			msg.Description.Default = alert.Annotations.Description
		}

		if len(alert.Labels) != 0 {
			msg.Labels = make(map[string]message.I10nField, len(alert.Labels))
			for k, v := range alert.Labels {
				msg.Labels[k] = message.I10nField{
					Default: v,
				}
				if strings.Compare("severity", k) == 0 {
					msg.Severity.Default = v
				}
			}
		}
		msg.Status.Default = alert.Status

		st, err := time.Parse(time.RFC3339Nano, alert.StartsAt)
		if err != nil {
			return []message.Message{}, err
		}
		msg.StartAt = st

		et, err := time.Parse(time.RFC3339Nano, alert.StartsAt)
		if err != nil {
			return []message.Message{}, err
		}
		msg.EndAt = et

		messages = append(messages, msg)
	}
	return messages, nil
}
