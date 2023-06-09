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

// CreateComment godoc
// @Summary      Create Comment
// @Description  Create a comment and store to database
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photo_id query int true "Photo ID"
// @Param        photo body models.InsertComment true "Field for insert comment"
// @Success      201  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /comments [post]
func (h HttpServer) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var req models.InsertComment
	userData := c.MustGet("userData").(jwt.MapClaims)
	var err error

	if contentType == "application/json" {
		err = c.ShouldBindJSON(&req)
	} else {
		err = c.ShouldBind(&req)
	}

	photoID := c.Query("photo_id")
	photoIDInt, er := strconv.Atoi(photoID)
	if er != nil {
		helpers.BadRequest(c, "Bad Query Parameter", er)
		return
	}
	req.UserID = uint(userData["id"].(float64))
	req.PhotoID = uint(photoIDInt)

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
			helpers.BadRequest(c, "Failed to Add New Comment", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	p, err := h.app.CreateComment(req)
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}
	helpers.Created(c, "Successfully Add Comment!", p)
}

// GetComments godoc
// @Summary      Get All Comment by Photo Id
// @Description  Get list of comments based on query params of photo id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        photo_id query int true "Photo ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /comments [get]
func (h HttpServer) GetComments(c *gin.Context) {
	photoID := c.Query("photo_id")
	if photoID == "" {
		helpers.BadRequest(c, "The photo id doesn't exist")
		return
	}

	photoIDInt, err := (strconv.Atoi(photoID))
	if err != nil {
		helpers.BadRequest(c, "Bad Query Parameter", err)
		return
	}

	res, err := h.app.GetComments(uint(photoIDInt))
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}

	helpers.OkWithData(c, "Success Retrive All Comments", res)
}

// GetCommentById godoc
// @Summary      Get Comment By Id
// @Description  Get a comment by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        commentId path string true "Comment ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /comments/{commentId} [get]
func (h HttpServer) GetCommentById(c *gin.Context) {
	req, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, err := h.app.GetCommentById(int64(req))
	if err != nil {
		if err.Error() == "record not found" {
			helpers.NotFound(c, "Comment Not Found")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
	}

	helpers.OkWithData(c, "Success Retrive A Comment", res)
}

// UpdateComment godoc
// @Summary      Update Comment
// @Description  Update a comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        commentId path string true "Comment Id"
// @Param        comment body models.UpdateComment true "Update a Comment"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /comments/{commentId} [put]
func (h HttpServer) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var comment models.UpdateComment
	userData := c.MustGet("userData").(jwt.MapClaims)
	var err error
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	comment.UserID = uint(userData["id"].(float64))
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
			helpers.BadRequest(c, "Failed to Update Comment!", res)
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}

	param, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}
	res, err := h.app.UpdateComment(int64(param), comment)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == "record not found" {
			helpers.NotFound(c, "Comment Not Found")
			return
		} else if err.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to update this Comment")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}
	helpers.OkWithData(c, "Successfully Updated Comment!", res)
}

// DeleteComment godoc
// @Summary      Delete Comment
// @Description  Delete Comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        commentId path string true "Comment ID"
// @Success      200  {object}  helpers.Response
// @Failure      400  {object}  helpers.Response
// @Failure      401  {object}  helpers.Response
// @Failure      404  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /comments/{commentId} [delete]
func (h HttpServer) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	req, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, er := h.app.DeleteComment(int64(req), userID)
	if er != nil {
		if er.Error() == "record not found" {
			helpers.NotFound(c, "Comment Not Found")
			return
		} else if er.Error() == "unauthorized" {
			helpers.BadRequest(c, "You are not authorized to delete this Comment")
			return
		}
		helpers.BadRequest(c, "Bad Request", err)
		return
	}
	helpers.OkWithData(c, "Comment Deleted Successfully!", res)
}
