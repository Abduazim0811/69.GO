package models


type Task struct{
	ID    		int32  		`json:"id"`
	Title  		string 		`json:"title"`
	Description string 		`json:"description"`
	Done 		bool  		`json:"done"` 
}