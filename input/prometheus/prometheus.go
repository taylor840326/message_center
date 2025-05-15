package prometheus

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
