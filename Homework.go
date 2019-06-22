package main	
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
)

type Todo struct {
	ID int `json:"id"`
	Title string  `json:"title"`
	Status string `json:"status"`
}

var Todos = map[int]Todo{}

func postTodoItem(c *gin.Context){
	s := Todo{}
	fmt.Printf("before bind %#v\n", s)
	if err := c.ShouldBindJSON(&s) ; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("after bind % #v\n" , s)

	id := len(Todos)
	id++
	s.ID = id
    Todos[id] = s
	c.JSON(http.StatusCreated, s)
}

func getListTodoItem(c *gin.Context){
	todoSlice := []Todo{}
	for _, list := range Todos {
		todoSlice = append(todoSlice, list)
	}
	c.JSON(http.StatusOK , todoSlice)
}

func getTodoItem(c *gin.Context){
	id := c.Param("id")
	idint, pass := strconv.Atoi(id)
	if pass != nil {
		c.JSON(http.StatusBadRequest, pass.Error())
		return
	}
	item , ok := Todos[idint]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, item)
}

func updateTodoItem(c *gin.Context){
	id := c.Param("id")
	idint, ok := strconv.Atoi(id)
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	item := Todos[idint]
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	Todos[idint] = item
	c.JSON(http.StatusOK, item)
}
func deleteTodoItem(c *gin.Context){
	id := c.Param("id")
	idint, ok := strconv.Atoi(id)
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	delete(Todos, idint)
	c.JSON(http.StatusOK, "success")
}
func main() {
	r := gin.Default()
	r.POST("/api/todos", postTodoItem)
	r.GET("/api/todos", getListTodoItem)
	r.GET("/api/todos/:id", getTodoItem)
	r.PUT("/api/todos/:id", updateTodoItem)
	r.DELETE("/api/todos/:id", deleteTodoItem )
	
	r.Run(":1234")
}