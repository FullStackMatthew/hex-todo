package interfaces

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"example.com/hex-todo/backend/domain"
	"example.com/hex-todo/backend/usecases"
)

type Handlers struct {
	TodoInteractor usecases.TodoInteractor
}

func (h *Handlers) AddTodoHandler(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.TodoInteractor.AddTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handlers) GetTodoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.TodoInteractor.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handlers) GetAllTodosHandler(c *gin.Context) {
	todos, err := h.TodoInteractor.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *Handlers) UpdateTodoHandler(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.TodoInteractor.UpdateTodo(todo.ID, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handlers) DeleteTodoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.TodoInteractor.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}
