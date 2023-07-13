package sqlstore

import "touchon_home/internal/model"

type HomeServerRepository struct {
	store *Store
}

// GetByUser получение данных о клиентских серверах, которые привязаны к аккаунту
func (serverRep *HomeServerRepository) GetByUser(userID int) ([]model.HomeServer, error) {
	var server model.HomeServer
	var servers []model.HomeServer

	if err := serverRep.store.db.QueryRow(
		"SELECT id, name, local_server, remote_server, local_sockets_port, remote_sockets_port, local_api_port, remote_api_port, main_object WHERE `user`=?",
		userID).Scan(&server.Id, &server.Name, &server.LocalServer, &server.RemoteServer,
		&server.LocalSocketsPort, &server.RemoteSocketsPort, &server.LocalAPIPort, &server.RemoteAPIPort, &server.MainObject); err != nil {
		return nil, err
	}

	servers = append(servers, server)

	return servers, nil
}
