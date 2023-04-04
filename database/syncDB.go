package database

import "grpc-auth-micro/model"

func SyncDB() {
	DB.AutoMigrate(
		&model.User{},
	)
}