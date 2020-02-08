package main
import(
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type employee struct{
	gorm.Model
	first string `form:"first" binding:"required" gorm:"type:varchar(100)"`
	last string `form:"last" binding:"required" gorm:"type:varchar(100)"`
	contact string	`form:"contact" binding:"required"`
	address string	`form:"address" binding:"required"`
	salary string `form:"salary" binding:"required"`
	deptid string	`form:"dept" binding:"required"`
}
func main(){
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	 r.GET("/insert", GetPeople)
	 r.POST("/insert", record)
 	r.Run(":8080")

}
func GetPeople(c *gin.Context){
	c.HTML(200, "form.html", nil)
}
func record(c *gin.Context){
	db,err := gorm.Open("sqlite3","form.db")
	if err!=nil{
		fmt.Println("-------Error-------")
		panic("failed to connect database")
	}else{
		fmt.Println("data created")
	}
	defer db.Close()
	db.AutoMigrate(&employee{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&employee{})
	var e employee 
	first := c.PostForm("first")
	e.last = c.PostForm("last")
	e.contact = c.PostForm("contact")
	e.address = c.PostForm("address")
	e.deptid = c.PostForm("dept")
	e.salary = c.PostForm("salary")
	db.Create(&employee{first: first})	
}