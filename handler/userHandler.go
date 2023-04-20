package handler

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iqbaludinm/mygram-api/helpers"
	"github.com/iqbaludinm/mygram-api/models"
)

// Registre godoc
// @Summary      User Register to MyGram
// @Description  Register user in MyGram Service
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body models.RegisterUser true "Credential for Register"
// @Success      201  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /auth/register [post]
func (h HttpServer) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.RegisterUser
	var err error
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		log.Println(err)
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]helpers.Form, len(verr))
			for i, fe := range verr {
				res[i] = helpers.Form{
					Field:   fe.Field(),
					Message: helpers.FormValidationError(fe),
				}
			}
			log.Println(verr)
			helpers.BadRequest(c, "Failed to Add New User", res)
			return
		}
	}

	u, err := h.app.RegisterUser(user)
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}
	helpers.Created(c, "Register is successfully!", u)
}

// Login godoc
// @Summary      User Login to MyGram
// @Description  Logs in the user with the provided email and password.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body models.LoginUser true "Credential for Login"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /auth/login [post]
func (h HttpServer) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.LoginUser
	var err error

	if contentType == "application/json" {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		log.Println(err)
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]helpers.Form, len(verr))
			for i, fe := range verr {
				res[i] = helpers.Form{
					Field:   fe.Field(),
					Message: helpers.FormValidationError(fe),
				}
			}
			log.Println(verr)
			helpers.BadRequest(c, "Please enter valid data!", res)
			return
		}
	}
	u, err := h.app.LoginUser(user)
	if err != nil {
		log.Println(err)
		helpers.InternalServerError(c, err.Error())
		return
	}
	
	helpers.OkWithData(c, "Successfully Login!", u)
}