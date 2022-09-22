package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gic-vikash/mongodb_rest_api_golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session)*UserController {
	return &UserController{s}
}

func(uc UserController) CreateUser(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	u:=models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u);err!=nil{
	log.Println(err)
	}
	u.Id = bson.NewObjectId()
	if err := uc.session.DB("mongo-golang").C("users").Insert(u) ; err!=nil{
		w.WriteHeader(404)
	}
	uj,err:=json.MarshalIndent(u,"","")
	if err!=nil{
		log.Println(err)
	}
	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(201)
	fmt.Fprintln(w,uj)
}
func(uc UserController) GetUser(w http.ResponseWriter , r *http.Request , p httprouter.Params){
	id:=p.ByName("id") 
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u) ; err!=nil{
		w.WriteHeader(404)
	}
	
	jsonuser,err:=json.MarshalIndent(u,"","")
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w,jsonuser)
}

func(uc UserController) DeleteUser(w http.ResponseWriter , r *http.Request ,p httprouter.Params){
	id:=p.ByName("id") 
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("mongo-golang").C("users").Remove(oid) ; err!=nil{
		w.WriteHeader(404)
	}
	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,"user deletion is successfull")
}

// func(uc UserController) UpdateUser(w http.ResponseWriter , r *http.Request){

// }