package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"skillfactory_task_31.3.1/internal/models"
)

func (api *API) createPostPG(ctx *gin.Context) {
	var post models.Post

	if err := ctx.BindJSON(&post); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := api.repository.PostsPG.AddPost(post)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Status": "OK"})
}

func (api *API) getPostsPG(ctx *gin.Context) {
	posts, err := api.repository.PostsPG.Posts()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (api *API) updatePostPG(ctx *gin.Context) {
	var post models.UpdatePost

	if err := ctx.BindJSON(&post); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invailed request body")
		return
	}

	err := api.repository.PostsPG.UpdatePost(post)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (api *API) deletPostByIdPG(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invailed id parameter")
		return
	}

	err = api.repository.PostsPG.DeletePost(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)

}
