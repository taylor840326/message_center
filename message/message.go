package message

import "time"

//	{
//		"status": "firing",
//		"labels": {
//			"alertname": "MysqlDown",
//			"app_kubernetes_io_component": "mysql",
//			"app_kubernetes_io_instance": "alder38",
//			"app_kubernetes_io_managed_by": "kubeblocks",
//			"app_kubernetes_io_name": "apecloud-mysql",
//			"apps_kubeblocks_io_component_name": "mysql",
//			"instance": "10.42.1.76:9104",
//			"job": "kubeblocks-service",
//			"namespace": "default",
//			"node": "ceph02.dev1.lab",
//			"pod": "alder38-mysql-0",
//			"service": "alder38-mysql-headless",
//			"severity": "critical"
//		},
//		"annotations": {
//			"description": "MySQL is down. (instance: alder38-mysql-0)",
//			"summary": "MySQL is down"
//		},
//		"startsAt": "2025-05-13T07:50:09.124Z",
//		"endsAt": "0001-01-01T00:00:00Z",
//		"generatorURL": "http://kb-addon-prometheus-server-0:9090/graph?g0.expr=max_over_time%28mysql_up%5B1m%5D%29+%3D%3D+0\u0026g0.tab=1",
//		"fingerprint": "faf2ad6b966c3a53"
//	}

type I10nField struct {
	// default content
	Default string `json:"default"`
	// translation content:
	// zh_CN: xxx
	// en_US: xxx
	// ...
	Translations map[string]string `json:"value"`
}

type Message struct {
	// Summary contains the brief overview of the message
	Summary I10nField `json:"summary"`
	// Description contains the detailed explanation of the message
	Description I10nField `json:"description"`
	// Labels stores key-value pairs of metadata about the message
	Labels map[string]I10nField `json:"labels"`
	// Status indicates the current state of the message (e.g. firing, resolved)
	Status I10nField `json:"status"`
	// Severity represents the urgency/importance level of the message (e.g. critical, warning, info)
	Severity I10nField `json:"severity"`
	// StartAt is the timestamp when the message was first created/triggered
	StartAt time.Time `json:"start_at"`
	// EndAt is the timestamp when the message was resolved/ended
	EndAt time.Time `json:"end_at"`
}
