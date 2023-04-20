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

// CreatePhoto godoc
// @Summary      Create Photo
// @Description  Store a photo to database
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photo body models.InsertPhoto true "Field for insert photo"
// @Success      201  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /photos [post]
func (h HttpServer) CreatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var req models.InsertPhoto
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
			helpers.BadRequest(c, "Failed to Add New Photo", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	p, err := h.app.CreatePhoto(req)
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}
	helpers.Created(c, "Successfully Add Photo!", p)
}

// GetPhotos godoc
// @Summary      Get All Photos
// @Description  Get list of photos with data relation (user)
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /photos [get]
func (h HttpServer) GetPhotos(c *gin.Context) {
	res, err := h.app.GetPhotos()
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}

	helpers.OkWithData(c, "Success Retrive All Photos", res)
}

// GetPhotoById godoc
// @Summary      Get Photo by Id
// @Description  Get details of photo corresponding with input id
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photoId path string true "Photo ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /photos/{photoId} [get]
func (h HttpServer) GetPhotoById(c *gin.Context) {
	req, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, err := h.app.GetPhotoById(int64(req))
	if err != nil {
		if err.Error() == "record not found" {
			helpers.NotFound(c, "Photo Not Found")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
	}

	helpers.OkWithData(c, "Success Retrive A Photo", res)
}

// UpdatePhoto godoc
// @Summary      Update Photo
// @Description  Update a photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photoId path string true "Photo Id"
// @Param        photo body models.UpdatePhoto true "Update a photo"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /photos/{photoId} [put]
func (h HttpServer) UpdatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var photo models.UpdatePhoto
	userData := c.MustGet("userData").(jwt.MapClaims)
	var err error
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&photo)
	} else {
		err = c.ShouldBind(&photo)
	}

	photo.UserID = uint(userData["id"].(float64))
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
			helpers.BadRequest(c, "Failed to Update Photo!", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	param, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}
	res, err := h.app.UpdatePhoto(int64(param), photo)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == "record not found" {
			helpers.NotFound(c, "Photo Not Found")
			return
		} else if err.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to update this Photo")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	helpers.OkWithData(c, "Successfully Updated Photo!", res)
}

// DeletePhoto godoc
// @Summary      Delete Photo
// @Description  Delete a photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photoId path string true "Photo ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /photos/{photoId} [delete]
func (h HttpServer) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	req, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, er := h.app.DeletePhoto(int64(req), userID)
	if er != nil {
		if er.Error() == "record not found" {
			helpers.NotFound(c, "Photo Not Found")
			return
		} else if er.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to delete this Photo")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}
	helpers.OkWithData(c, "Photo Deleted Successfully!", res)
}
