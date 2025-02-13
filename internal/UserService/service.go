package UserService

import "awesomeProject1/internal/taskService"

type UserService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) DeleteUserByID(id string) error {
	return s.repo.DeleteUserByID(id)
}
func (s *UserService) PatchUserByID(id uint, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)

}
func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return s.repo.GetUsersUserId(userID)
}
