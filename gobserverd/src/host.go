package src

type Service struct {
	Port    int
	Name    string
	Version string
}

type Host struct {
	IP       string
	Services []Service
}

func NewHost(ip string) *Host {
	var h Host
	h.IP = ip
	return &h
}
