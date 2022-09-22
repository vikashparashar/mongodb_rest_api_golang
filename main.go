// 23/09/2022 Vikash Parashar
// golang_mongo_rest_api's

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gic-vikash/mongodb_rest_api_golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("starting the application")
	r := httprouter.New()
	uc:=controllers.NewUserController(getSession())
	r.POST("/user",uc.CreateUser)
	r.GET("/user/:id",uc.GetUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	// r.PUT("/user/:id",uc.UpdateUser)
	if err := http.ListenAndServe(":8888",r) ; err!=nil{
		log.Fatal(err)
	}
	fmt.Println("server is up and running on port : 8888 on your local machine")
}


func getSession()*mgo.Session{
	s,err:=mgo.Dial("mongodb://localhost:27017")
	if err!=nil{
		log.Fatal(err)
	}
	return s
}
