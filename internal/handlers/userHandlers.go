package handlers

import (
	"awesomeProject1/internal/UserService" // Импортируем наш сервис
	"awesomeProject1/internal/web/Users"
	"golang.org/x/net/context"
)

type UserHandler struct {
	Service *UserService.UserService
}

func NewUserHendler(service *UserService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (handler *UserHandler) GetUsers(ctx context.Context, request Users.GetUsersRequestObject) (Users.GetUsersResponseObject, error) {
	allUsers, err := handler.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := Users.GetUsers200JSONResponse{}
	for _, user := range allUsers {
		User := Users.User{
			Id:       &user.Id,
			Email:    &user.Email,
			Password: &user.Password,
		}
		response = append(response, User)
	}
	return response, nil
}
func (Handler *UserHandler) PostUsers(ctx context.Context, request Users.PostUsersRequestObject) (Users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := UserService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := Handler.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := Users.PostUsers201JSONResponse{
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, err
}

func (Handler *UserHandler) DeleteUsersId(ctx context.Context, request Users.DeleteUsersIdRequestObject) (Users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := Handler.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}
	response := Users.DeleteUsersId204Response{}
	return response, nil
}
func (Handler *UserHandler) PatchUsersId(ctx context.Context, request Users.PatchUsersIdRequestObject) (Users.PatchUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body
	userToUpdate := UserService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	updatedUser, err := Handler.Service.PatchUserByID(id, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := Users.PatchUsersId200JSONResponse{
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
		Id:       &updatedUser.Id,
	}
	return response, nil

}
