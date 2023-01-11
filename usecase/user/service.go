package user

type UserService interface {
	Register(req RegisterRequest) (res RegisterResponse, err error)
	Login(req LoginRequest) (res LoginResponse, err error)
	User(cookies string) (res UserResponse, err error)
	Logout() (res LogoutResponse, err error)
}
