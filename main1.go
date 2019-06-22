package main
import (
	   "net/http"
	   "github.com/gin-gonic/gin"
	   "fmt"
	  
)
type Student struct {
	//Name string
	//ID string
	Name string `json:"name"`
	ID   int `json:"student_id"`
}
func pingHandler (c *gin.Context){
	//c.JSON(200,gin.H{
	response := gin.H{
		"message": 1,
	}
	c.JSON(http.StatusOK, response)
}
func pingPostHandler (c *gin.Context){
	response := gin.H{
		"message": "this is post",
	}
	c.JSON(http.StatusOK, response)
}

var ss = map[int]Student {
	1: Student {Name: "Anuchit" , ID : 1},
}

func getStudentHander (c *gin.Context) {
		//ss := map[string]Student {
		//"965108": Student {Name: "Anuchit" , ID : "965108"},
		//}
	//c.JSON(http.StatusOK, ss)
	studentSlice := []Student{}
	for _,v := range ss {
		//fmt.Printf("key[%s] value[%s]\n", k, ss[k])
		studentSlice = append(studentSlice , v)
	}
	//st := []Student{Student{Name:"111"},Student{Name:"222"}}
	c.JSON(http.StatusOK, studentSlice)
}

func postStudentHander (c *gin.Context) {
	//receive -> student {...}
	//add student -> map ss
	s := Student{}
	fmt.Printf("before bind % #v\n" , s)
	if err := c.ShouldBindJSON(&s) ; err != nil {
		//c.JSON(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("after bind % #v\n" , s)
	
	id := len(ss)
	id++
	s.ID = id
	ss[id] = s
	c.JSON(http.StatusOK, s)
}
func main(){
	r := gin.Default()

	r.GET("/ping", pingHandler) 
	r.POST("/ping",pingPostHandler )
	r.GET("/students", getStudentHander) 
	r.POST("/students", postStudentHander) 
		
//r.Run()
r.Run(":1234")
}