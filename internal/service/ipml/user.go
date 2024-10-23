package ipml

import "fmt"

type UserSrvImpl struct {
}

func NewUserSrvImpl() *UserSrvImpl {
	return &UserSrvImpl{}
}

func (u UserSrvImpl) Login() {
	fmt.Println("user login success")
}

func (u UserSrvImpl) Logout() {
	fmt.Println("user logout success")
}

func (u UserSrvImpl) Register() {
	fmt.Println("user register success")
}
