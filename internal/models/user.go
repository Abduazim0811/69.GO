package models


type Task struct{
	ID    		int32  		`json:"id"`
	Title  		string 		`json:"title"`
	Description string 		`json:"description"`
	Done 		bool  		`json:"done"` 
}

type StandartError struct{
	Error error `json:"error"`
}

type ForbiddenError struct{
	Message string `json:"message"`
}