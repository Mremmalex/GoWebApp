package main 

import (

	"fmt"
	"net/http"
	"Webtutorial/routes"
	"Webtutorial/utils"
	"Webtutorial/models"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r :=routes.NewRouter()
	 
	http.Handle("/", r)
	fmt.Println("servering application at port :8080")
	http.ListenAndServe(":8080", nil)
}
