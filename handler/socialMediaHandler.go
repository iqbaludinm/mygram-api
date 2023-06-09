package handler

import (
	"errors"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iqbaludinm/mygram-api/helpers"
	"github.com/iqbaludinm/mygram-api/models"
)

// CreateSocialMedia godoc
// @Summary      Create Social Media
// @Description  Add Social Media in User Data
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        socmed body models.InsertSocialMedia true "Field for insert social media"
// @Success      201  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /social-medias [post]
func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var req models.InsertSocialMedia
	userData := c.MustGet("userData").(jwt.MapClaims)
	var err error

	if contentType == "application/json" {
		err = c.ShouldBindJSON(&req)
	} else {
		err = c.ShouldBind(&req)
	}

	req.UserID = uint(userData["id"].(float64))
	if err != nil {
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
			helpers.BadRequest(c, "Failed to Add New Socmed", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	p, err := h.app.CreateSocmed(req)
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}
	helpers.Created(c, "Successfully Add Social Media!", p)
}

// GetSocialMedias godoc
// @Summary      Get All Social Media
// @Description  Get list of social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /social-medias [get]
func (h HttpServer) GetSocialMedias(c *gin.Context) {
	res, err := h.app.GetSocmeds()
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}

	helpers.OkWithData(c, "Success Retrive All Social Media", res)
}

// GetSocialMediaById godoc
// @Summary      Get Social Media by Id
// @Description  Get details of social media corresponding with social media id
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        socmedId path string true "Social Media ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /social-medias/{socmedId} [get]
func (h HttpServer) GetSocialMediaById(c *gin.Context) {
	req, err := strconv.Atoi(c.Param("socmedId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, err := h.app.GetSocmedById(int64(req))
	if err != nil {
		if err.Error() == "record not found" {
			helpers.NotFound(c, "The Social Media Not Found")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
	}

	helpers.OkWithData(c, "Success Retrive A Social Media Profile", res)
}

// UpdateSocialMedia godoc
// @Summary      Update Social Media
// @Description  Update a social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        socmedId path string true "Social Media Id"
// @Param        socmed body models.UpdateSocialMedia true "Update a social media"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /social-medias/{socmedId} [put]
func (h HttpServer) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var socmed models.UpdateSocialMedia
	userData := c.MustGet("userData").(jwt.MapClaims)
	var err error
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&socmed)
	} else {
		err = c.ShouldBind(&socmed)
	}

	socmed.UserID = uint(userData["id"].(float64))
	if err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]helpers.Form, len(verr))
			for i, fe := range verr {
				res[i] = helpers.Form{
					Field:   fe.Field(),
					Message: helpers.FormValidationError(fe),
				}
			}
			helpers.BadRequest(c, "Failed to Update Social Media Data!", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	param, err := strconv.Atoi(c.Param("socmedId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}
	res, err := h.app.UpdateSocmed(int64(param), socmed)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == "record not found" {
			helpers.NotFound(c, "Social Media Not Found")
			return
		} else if err.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to update this Social Media Data")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}
	helpers.OkWithData(c, "Successfully Updated Socmed!", res)
}

// DeleteSocialMedia godoc
// @Summary      Delete Social Media
// @Description  Delete a Social Media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        socmedId path string true "Social Media ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /social-medias/{socmedId} [delete]
func (h HttpServer) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	req, err := strconv.Atoi(c.Param("socmedId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, er := h.app.DeleteSocmed(int64(req), userID)
	if er != nil {
		if er.Error() == "record not found" {
			helpers.NotFound(c, "Social Media Not Found")
			return
		} else if er.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to delete this Social Media")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}
	helpers.OkWithData(c, "Socmed Deleted Successfully!", res)
}
