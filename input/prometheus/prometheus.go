package prometheus

//{
// "status": "firing",
// "labels": {
// 	"alertname": "MysqlDown",
// 	"app_kubernetes_io_component": "mysql",
// 	"app_kubernetes_io_instance": "alder38",
// 	"app_kubernetes_io_managed_by": "kubeblocks",
// 	"app_kubernetes_io_name": "apecloud-mysql",
// 	"apps_kubeblocks_io_component_name": "mysql",
// 	"instance": "10.42.1.76:9104",
// 	"job": "kubeblocks-service",
// 	"namespace": "default",
// 	"node": "ceph02.dev1.lab",
// 	"pod": "alder38-mysql-0",
// 	"service": "alder38-mysql-headless",
// 	"severity": "critical"
// },
// "annotations": {
// 	"description": "MySQL is down. (instance: alder38-mysql-0)",
// 	"summary": "MySQL is down"
// },
// "startsAt": "2025-05-13T07:50:09.124Z",
// "endsAt": "0001-01-01T00:00:00Z",
// "generatorURL": "http://kb-addon-prometheus-server-0:9090/graph?g0.expr=max_over_time%28mysql_up%5B1m%5D%29+%3D%3D+0\u0026g0.tab=1",
// "fingerprint": "faf2ad6b966c3a53"
// }

// labels
type Labels struct {
	AlertName                     string `json:"alertname"`
	AppKubernetesIOComponent      string `json:"app_kubernetes_io_component"`
	AppKubernetesIOInstance       string `json:"app_kubernetes_io_instance"`
	AppKubernetesIOManagedBy      string `json:"app_kubernetes_io_managed_by"`
	AppKubernetesIOName           string `json:"app_kubernetes_io_name"`
	AppsKubeblocksIOComponentName string `json:"apps_kubeblocks_io_component_name"`
	Instance                      string `json:"instance"`
	Job                           string `json:"job"`
	Namespace                     string `json:"namespace"`
	Node                          string `json:"node"`
	Pod                           string `json:"pod"`
	Service                       string `json:"service"`
	Severity                      string `json:"severity"`
}

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

type AlertMessage struct {
	Alerts []AlertItem `json:"alerts"`
}
