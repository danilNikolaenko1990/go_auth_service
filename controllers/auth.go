package controllers

import (
	"auth-service/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	ContentTypeJson       = "application/json"
	ContentTypeHeaderName = "Content-Type"
)

type AuthController struct {
	loginChecker service.LoginChecker
	userCreator  service.UserCreator
	validator    *validator.Validate
}

func NewAuthController(loginChecker service.LoginChecker,
	userCreator service.UserCreator, validator *validator.Validate) AuthController {

	return AuthController{
		loginChecker: loginChecker,
		userCreator:  userCreator,
		validator:    validator,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginAction LoginAction
	err := decoder.Decode(&loginAction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.validator.Struct(loginAction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isLogged, err := c.loginChecker.IsLoggedIn(loginAction.Login, loginAction.Password)
	resp := IsLoggedInResponse{}
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.sendRespToClient(w, err, resp)
		return
	}

	resp.IsLogged = isLogged
	c.sendRespToClient(w, err, resp)
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var createUser CreateUserAction
	err := decoder.Decode(&createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.validator.Struct(createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := c.createUser(createUser)
	userCreationResult := c.convertCreationResp(resp)
	c.sendRespToClient(w, err, userCreationResult)
}

func (c *AuthController) sendRespToClient(w http.ResponseWriter, err error, userCreationResult interface{}) {
	respInJson, err := json.Marshal(userCreationResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(ContentTypeHeaderName, ContentTypeJson)
	w.Write(respInJson)
}

func (c *AuthController) convertCreationResp(resp service.CreationResponse) UserCreationResult {
	userCreationResult := UserCreationResult{
		Registered:   resp.Registered,
		ErrorMessage: resp.ErrorMessage,
		ErrorCode:    resp.ErrorCode,
	}
	return userCreationResult
}

func (c *AuthController) createUser(createUser CreateUserAction) service.CreationResponse {
	resp := c.userCreator.Create(service.CreationRequest{
		Login:    createUser.Login,
		Email:    createUser.Email,
		Phone:    createUser.Phone,
		Password: createUser.Password,
	})
	return resp
}
