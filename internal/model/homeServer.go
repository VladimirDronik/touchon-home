package model

type ServersStruct struct {
	Servers []HomeServer
}

type HomeServer struct {
	Id                int
	Name              string
	LocalServer       string
	RemoteServer      string
	LocalSocketsPort  int
	RemoteSocketsPort int
	LocalAPIPort      int
	RemoteAPIPort     int
	MainObject        bool
}
