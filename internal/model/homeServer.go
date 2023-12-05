package model

type ServersStruct struct {
	Servers []HomeServer
}

type HomeServer struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	LocalServer       string `json:"local-server"`
	RemoteServer      string `json:"remote-server"`
	LocalSocketsPort  int    `json:"local-sockets-port"`
	RemoteSocketsPort int    `json:"remote-sockets-port"`
	LocalAPIPort      int    `json:"local-API-port"`
	RemoteAPIPort     int    `json:"remote-API-port"`
	LocalCCTVPort     int    `json:"local-CCTV-port"`
	RemoteCCTVPort    int    `json:"remote-CCTV-port"`
	MainObject        bool   `json:"main-object"`
}
