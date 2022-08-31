package main

import (
	"Latihan_1/app/domain/dto"
	"Latihan_1/app/internal/config"
	repository "Latihan_1/app/internal/repository/psql/user"
	"Latihan_1/app/internal/repository/psql/user/model"
	"Latihan_1/app/internal/usecase/user"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Latihan 1")
	// Koneksi ke db
	db, err := config.ConnectPostgreSQL()
	if err != nil {
		log.Fatal("Gagal Koneksi DB :", err.Error())
	}

	fmt.Printf("Success Connect DB : %s\n", "latihan_1")

	// Migrasi Model
	db.AutoMigrate(&model.User{})

	// Panggil user repository
	userRepository := repository.NewUserRepository(db)
	userUsecase := user.NewUserUsecase(userRepository)

	newUserData := dto.User{
		Name:  "Hendra",
		Age:   21,
		Email: "hendra@gmail.com",
	}
	createdUser, err := userUsecase.CreateNewUser(&newUserData)

	if err != nil {
		log.Fatal("Ada Error Nih", err.Error())
	}

	fmt.Println("Sukses Create New User :", createdUser)
}
