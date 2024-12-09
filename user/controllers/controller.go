package controllers

import serviceinterfaces "github.com/anandtiwari11/IEKart-go/user/serviceInterfaces"

type UserController struct {
	UserService serviceinterfaces.IUserServiceInterface
}

func NewUserController(userService serviceinterfaces.IUserServiceInterface) *UserController {
    return &UserController{
        UserService: userService,
    }
}
