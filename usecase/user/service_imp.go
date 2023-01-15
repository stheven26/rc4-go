package user

import (
	"hashing-file/constants"
	"hashing-file/domain/entity"
	"hashing-file/domain/repository"
	"hashing-file/helpers"
	"hashing-file/middlewares"
	"strconv"
	"time"
)

type service struct {
	userRepository repository.UserRepository
}

func NewService(user repository.UserRepository) UserService {
	return &service{userRepository: user}
}

func (s *service) Register(req RegisterRequest) (res RegisterResponse, err error) {
	hash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return
	}
	req.Password = hash
	data := entity.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.userRepository.Register(data)
	if err != nil {
		return
	}
	res = RegisterResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: RegisterData{
			Username: data.Username,
			Email:    data.Email,
		},
	}
	return
}

func (s *service) Login(req LoginRequest) (res LoginResponse, err error) {
	users, err := s.userRepository.Login(req.Email)
	if err != nil {
		return
	}
	if users.ID == 0 {
		return
	}
	if _, err = helpers.CheckHashPassword(req.Password, users.Password); err != nil {
		return
	}
	token, err := middlewares.CreateToken(strconv.Itoa(int(users.ID)))
	if err != nil {
		return
	}
	res = LoginResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: LoginData{
			Email: req.Email,
			Token: token,
		},
	}
	return
}

func (s *service) User(cookies string) (res UserResponse, err error) {
	token, err := middlewares.Auth(cookies)
	if err != nil {
		return
	}
	claims := middlewares.ExtractAuth(token)
	user, err := s.userRepository.User(claims.Issuer)
	if err != nil {
		return
	}
	res = UserResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: UserData{
			User:    user,
			Expires: claims.ExpiresAt,
			Time:    time.Now(),
		},
	}
	return
}
func (s *service) Logout() (res LogoutResponse, err error) {
	cookie, err := middlewares.ClearToken()
	if err != nil {
		return
	}
	res = LogoutResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: LogoutData{
			Cookies: cookie,
		},
	}
	return
}

func (s *service) GetAllUser() (res DefaultResponse, err error) {
	user, err := s.userRepository.GetAllUser()
	if err != nil {
		return
	}
	res = DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    user,
	}
	return
}
