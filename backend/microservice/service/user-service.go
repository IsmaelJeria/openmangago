package service

import (
	"direst/domain"
	"direst/repository"
	"encoding/json"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//UserService ..
type UserService struct {
	userRepository repository.IUserRepository
}

//NewUserService ..
func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{userRepository: repository}
}

//FindByID ..
func (s *UserService) FindByID(id string) (int, []byte) {
	d, err := s.userRepository.FindByID(id)
	var user domain.User
	if err != nil {
		log.Println("s.userRepository.FindByID(id", err)
		return 500, []byte(`Internal server Error`)
	}
	if d.Exists() == false {
		return 404, []byte(`Resource not found`)
	}

	b, err := json.Marshal(d.Data())
	err = json.Unmarshal(b, &user)

	if err != nil {
		log.Println("UserService json.Marshal(d.Data())", err)
		return 500, []byte(`Internal server Error`)
	}

	return 200, b
}

//Save ..
func (s UserService) Save(requestBody []byte) (int, []byte) {

	var user domain.User
	json.Unmarshal(requestBody, &user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword([]byte(user.Password)", err)
		return 500, []byte(`Internal server Error`)
	}
	user.Password = string(hashedPassword)
	s.userRepository.Save(user.Name, user)

	return 200, []byte(`user created`)
}
