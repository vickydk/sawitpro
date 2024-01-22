package handler

import (
	"fmt"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) createNewUser(request *generated.SignupRequest) (userUUID uuid.UUID, err error) {
	_, err = s.Repository.FindUserByPhoneNumber(request.PhoneNumber)
	if err == nil {
		err = fmt.Errorf("phone number already exists")
		return
	}
	entity := &repository.Users{
		PhoneNumber: request.PhoneNumber,
		Name:        request.FullName,
		Password:    hashAndSalt([]byte(request.Password)),
	}
	err = s.Repository.SaveUser(entity)
	if err != nil {
		return
	}
	userUUID = entity.ID
	return
}

func (s *Server) checkAuth(request *generated.LoginRequest) (resp *generated.LoginResponse, err error) {
	user, err := s.Repository.FindUserByPhoneNumber(request.PhoneNumber)
	if err != nil {
		err = fmt.Errorf("user or password not match")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		err = fmt.Errorf("user or password not match")
		return
	}
	resp = &generated.LoginResponse{
		UserUUID: user.ID,
	}
	resp.Token, _ = NewBearerToken(user)
	return
}

func (s *Server) findUserByID(id string) (resp *generated.UserResponse, err error) {
	user, err := s.Repository.FindUserById(id)
	if err != nil {
		err = fmt.Errorf("user not found")
		return
	}
	resp = &generated.UserResponse{
		FullName:    user.Name,
		PhoneNumber: user.PhoneNumber,
	}
	return
}

func (s *Server) userUpdate(id string, request *generated.UpdateUserRequest) (resp *generated.UserResponse, err error) {
	user, err := s.Repository.FindUserById(id)
	if err != nil {
		err = fmt.Errorf("user not found")
		return
	}
	if user.PhoneNumber != request.PhoneNumber {
		_, err = s.Repository.FindUserByPhoneNumber(request.PhoneNumber)
		if err == nil {
			err = fmt.Errorf("phone number already exists")
			return
		}
	}
	user.PhoneNumber = request.PhoneNumber
	user.Name = request.FullName
	user.UpdatedAt = time.Now()
	err = s.Repository.SaveUser(user)
	if err != nil {
		return
	}
	resp = &generated.UserResponse{
		FullName:    user.Name,
		PhoneNumber: user.PhoneNumber,
	}
	return
}
