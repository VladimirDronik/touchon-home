package model

type ServersStruct struct {
	Servers []HomeServer `json:"servers"`
}

type HomeServer struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	LocalServer       string `json:"localServer"`
	RemoteServer      string `json:"remoteServer"`
	LocalSocketsPort  int    `json:"localSocketsPort"`
	RemoteSocketsPort int    `json:"remoteSocketsPort"`
	LocalAPIPort      int    `json:"localAPIPort"`
	RemoteAPIPort     int    `json:"remoteAPIPort"`
	LocalCCTVPort     int    `json:"localCCTVPort"`
	RemoteCCTVPort    int    `json:"remoteCCTVPort"`
	MainObject        bool   `json:"mainObject"`
}
