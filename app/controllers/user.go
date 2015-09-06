/*
 * account.go
 */

package controllers

import (
	"github.com/revel/revel"
	//	"github.com/robfig/revel"
	//"goweb/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Signup() revel.Result {
	return c.Render()
}

//func (c User) SignupPost(user *models.MockUser) revel.Result {
//	return c.Render()
//}

func (c User) Signin() revel.Result {
	return c.Render()
}
