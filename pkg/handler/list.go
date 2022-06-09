package handler

import (
	"awesomeProject"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not found")
		return
	}

	var input awesomeProject.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newListId, err := h.services.TodoList.Create(id.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id":     newListId,
		"title":       input.Title,
		"description": input.Description,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not found")
	}

	lists, err := h.services.TodoList.GetAll(id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lists": lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "cannot convert query param")
	}

	list, err := h.services.TodoList.GetById(listId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id":     listId,
		"title":       list.Title,
		"description": list.Description,
	})
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.TodoList.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
