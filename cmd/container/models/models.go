package models

import "time"

type Container struct {
	Id              string          `json:"id"`
	Image           string          `json:"image"`
	State           string          `json:"state"`
	Status          string          `json:"status"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
}

type bridge struct {
	Gateway    string `json:"Gateway"`
	IPAddress  string `json:"IPAddress"`
	MacAddress string `json:"MacAddress"`
}
type Networks struct {
	Bridge bridge `json:"bridge"`
}
type NetworkSettings struct {
	Networks Networks `json:"Networks"`
}

type State struct {
	Status     string    `json:"Status"`
	Running    bool      `json:"Running"`
	Paused     bool      `json:"Paused"`
	Restarting bool      `json:"Restarting"`
	OOMKilled  bool      `json:"OOMKilled"`
	Dead       bool      `json:"Dead"`
	Pid        int       `json:"Pid"`
	ExitCode   int       `json:"ExitCode"`
	Error      string    `json:"Error"`
	StartedAt  time.Time `json:"StartedAt"`
	FinishedAt time.Time `json:"FinishedAt"`
}

type HostConfig struct {
	NetworkMode        string `json:"NetworkMode"`
	Cgroup             string `json:"Cgroup"`
	PidMode            string `json:"PidMode"`
	Privileged         bool   `json:"Privileged"`
	PublishAllPorts    bool   `json:"PublishAllPorts"`
	ReadonlyRootfs     bool   `json:"ReadonlyRootfs"`
	SecurityOpt        any    `json:"SecurityOpt"`
	UTSMode            string `json:"UTSMode"`
	UsernsMode         string `json:"UsernsMode"`
	ShmSize            int    `json:"ShmSize"`
	Runtime            string `json:"Runtime"`
	Isolation          string `json:"Isolation"`
	CPUShares          int    `json:"CpuShares"`
	Memory             int    `json:"Memory"`
	NanoCpus           int    `json:"NanoCpus"`
	CgroupParent       string `json:"CgroupParent"`
	CPUPeriod          int    `json:"CpuPeriod"`
	CPUQuota           int    `json:"CpuQuota"`
	CPURealtimePeriod  int    `json:"CpuRealtimePeriod"`
	CPURealtimeRuntime int    `json:"CpuRealtimeRuntime"`
	CpusetCpus         string `json:"CpusetCpus"`
	CpusetMems         string `json:"CpusetMems"`
	DeviceCgroupRules  any    `json:"DeviceCgroupRules"`
	DeviceRequests     any    `json:"DeviceRequests"`
	MemoryReservation  int    `json:"MemoryReservation"`
	MemorySwap         int    `json:"MemorySwap"`
	MemorySwappiness   any    `json:"MemorySwappiness"`
	Ulimits            any    `json:"Ulimits"`
	CPUCount           int    `json:"CpuCount"`
	CPUPercent         int    `json:"CpuPercent"`
	IOMaximumIOps      int    `json:"IOMaximumIOps"`
	IOMaximumBandwidth int    `json:"IOMaximumBandwidth"`
}
type Config struct {
	Hostname     string `json:"Hostname"`
	Domainname   string `json:"Domainname"`
	User         string `json:"User"`
	AttachStdin  bool   `json:"AttachStdin"`
	AttachStdout bool   `json:"AttachStdout"`
	AttachStderr bool   `json:"AttachStderr"`
	Tty          bool   `json:"Tty"`
	OpenStdin    bool   `json:"OpenStdin"`
	StdinOnce    bool   `json:"StdinOnce"`
	Image        string `json:"Image"`
	WorkingDir   string `json:"WorkingDir"`
	OnBuild      any    `json:"OnBuild"`
}
type Inspect struct {
	Id              string          `json:"id"`
	State           State           `json:"State"`
	HostConfig      HostConfig      `json:"HostConfig"`
	Status          string          `json:"status"`
	Config          Config          `json:"Config"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
}
