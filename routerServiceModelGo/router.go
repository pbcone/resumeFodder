package main

import "github.com/julienschmidt/httprouter"

func filmPacRouter() *httprouter.Router {
	router := httprouter.New()
	//fpModel Routes
	router.POST("/fpmodel/create", CreateFpModel)
	router.GET("/fpmodel/all", GetAllFPModel)

	return router
}
