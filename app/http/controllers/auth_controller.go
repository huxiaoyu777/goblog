package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {

}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request)  {

	_user := user.User{
		Name: r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User": _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败")
		}
	}

}