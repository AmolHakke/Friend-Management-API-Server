package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	 "github.com/gorilla/mux"
	 "io/ioutil"
)

// type FirendList struct {
	
//     EmailID string `json:"EmailID"`
// 	FrndEmailID string `json:"FrndEmailID"`
   
    
// }


// type PersonList struct {
	
//     EmailID string `json:"EmailID"`
   
    
// }


type FirendList struct {
	
    EmailID string `json:"EmailID"`
	FrndEmailID string `json:"FrndEmailID"`   
    
}

 
var Firends []FirendList 

// var Persons []PersonList 

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnSingleArticle")
    	decoder := json.NewDecoder(r.Body)
	var data FirendList
	err := decoder.Decode(&data)
	if err != nil 	{
		panic(err)
	 }
	 fmt.Println(data.EmailID)
    for _, frned := range Firends {
        if frned.EmailID == data.EmailID {
			json.NewEncoder(w).Encode(frned.EmailID)
			json.NewEncoder(w).Encode("Success")
			fmt.Println(frned.EmailID)
			fmt.Println("Endpoint Hit: in if condition")
        }
    }
}
 

func create_NewFirend(w http.ResponseWriter, r *http.Request) {
	
	
	fmt.Println("Endpoint Hit: create_NewFirend")	
    reqBody, _ := ioutil.ReadAll(r.Body)
    var CreateNewFriend FirendList 
    json.Unmarshal(reqBody, &CreateNewFriend)
	// decoder := json.NewDecoder(r.Body)
	// var data FirendList
	// err := decoder.Decode(&data)
	// if err != nil 	{
	// 	panic(err)
	// }

	// fmt.Println("EmailID"+data.EmailID)
	// fmt.Println("FrndEmailID"+data.FrndEmailID)
    
	var isPersonValid bool =true
	
	if(isPersonValid)	{
		Firends = append(Firends, CreateNewFriend)
	 json.NewEncoder(w).Encode(CreateNewFriend)
	fmt.Println("Endpoint Hit: Success")
			 

	}
	 
// 	fmt.Println("Endpoint Hit: END")
   
} 



func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
   
	myRouter.HandleFunc("/CreateNewFriend", create_NewFirend).Methods("POST")
	myRouter.HandleFunc("/GetAllFriends", returnSingleArticle).Methods("POST")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}



func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Firends = []FirendList{
        FirendList{EmailID: "amol@gmail.com",FrndEmailID: "Frndamol1@gmail.com"},
        FirendList{EmailID: "amol@gmail.com",FrndEmailID: "Frndamol2@gmail.com"},
	}
	
// 	Persons = []PersonList{
//         PersonList{EmailID: "amol@gmail.com"},
//         PersonList{EmailID: "amol_1@gmail.com"},
//     }
    handleRequests()
}

