package store

import "touchon_home/internal/model"

type ServerRepository interface {
	GetByUser(userID int) ([]model.HomeServer, error) //Получение всех серверов пользователя
}
