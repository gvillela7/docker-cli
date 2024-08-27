package models

type Replicated struct {
	Replicas string `json:"replicas"`
}
type Mode struct {
	Replicated Replicated `json:"replicated"`
}
type Spec struct {
	Availability string `json:"availability"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	Mode         Mode   `json:"mode"`
}
type Status struct {
	State string `json:"state"`
	Addr  string `json:"addr"`
}
type Description struct {
	Hostname string `json:"hostname"`
}
type Ports struct {
	Protocol   string `json:"protocol"`
	TargetPort string `json:"targetPort"`
}
type Endpoint struct {
	Ports Ports `json:"ports"`
}
type Nodes struct {
	Id       string      `json:"id"`
	Hostname Description `json:"description"`
	Spec     Spec        `json:"spec"`
	State    string      `json:"state"`
	Status   Status      `json:"status"`
}

type Services struct {
	Id       string      `json:"id"`
	Hostname Description `json:"description"`
	Spec     Spec        `json:"spec"`
	Mode     Spec        `json:"mode"`
	Image    Status      `json:"status"`
	Ports    Endpoint    `json:"ports"`
}
