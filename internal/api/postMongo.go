package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"skillfactory_task_31.3.1/internal/models"
)

func (api *API) createPostMongo(ctx *gin.Context) {
	var post models.Post

	if err := ctx.BindJSON(&post); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := api.repository.PostsMongo.AddPost(post)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Status": "OK"})
}

func (api *API) getPostsMongo(ctx *gin.Context) {
	posts, err := api.repository.PostsMongo.Posts()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (api *API) updatePostMongo(ctx *gin.Context) {
	var post models.UpdatePost

	if err := ctx.BindJSON(&post); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invailed request body")
		return
	}

	err := api.repository.PostsMongo.UpdatePost(post)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (api *API) deletPostByIdMongo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invailed id parameter")
		return
	}

	err = api.repository.PostsMongo.DeletePost(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)

}
