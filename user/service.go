package user

type UserService interface {
	CreateUser(input CreateUserInput) (User, error)
	GetAll() ([]User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) CreateUser(input CreateUserInput) (User, error) {
	var newUser User

	newUser.Name = input.Name
	newUser.Email = input.Email
	newUser.Password = input.Password

	data, err := s.repository.Save(newUser)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *userService) GetAll() ([]User, error) {
	users, err := s.repository.FetchAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
