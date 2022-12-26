package iiko_transport

import (
	iikomodel "github.com/aschurf/iiko_transport/internal/models"
)

// NewIikoTransport
/*
	Создает новую авторизацю в IIKO transport api, возвращает ошибку если есть
*/
func NewIikoTransport(login string) (iikomodel.IikoLogin, error) {

	iikoTransport := iikomodel.IikoLogin{
		Login: login,
	}

	err := iikoTransport.Auth()
	if err != nil {
		return iikoTransport, err
	}

	return iikoTransport, nil
}
