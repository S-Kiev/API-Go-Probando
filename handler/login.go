package handler

import (
	//"encoding/json"
	"net/http"

	"github.com/S-Kiev/API-Go-Probando/authorization"
	"github.com/S-Kiev/API-Go-Probando/modelo"
	"github.com/labstack/echo"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := modelo.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "estructura no válida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña no válidos", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *modelo.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
