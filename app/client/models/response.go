package models

type Response struct {
	Status  int
	Message string
	Object  interface{}
}
