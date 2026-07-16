package main

import (
	"crudAPI/internal/database"
	"crudAPI/internal/handler"
	"crudAPI/internal/router"
	"fmt"
	"net/http"
	"runtime/debug"
)

func errorhandle(err error){
	if err != nil {
		panic(err)
	}
}
func panichandle(){
	r := recover()
	if r!=nil {
		fmt.Println("error : ",r)
		debug.PrintStack()
	}
}

func main() {

	defer panichandle()

	var newdb database.Usersdbstr
	newdb.Dbinit()

	newhandler := handler.NewHandler(&newdb)

	router.Router(newhandler)
	
	fmt.Println("server started at :4000")
	err:=http.ListenAndServe(":4000",nil)
	errorhandle(err)
	

	
}