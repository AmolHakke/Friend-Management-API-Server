//  package section
package main
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	 "github.com/gorilla/mux"
	 "io/ioutil"
	 "bytes"
	 "github.com/mcnijman/go-emailaddress"
	 	 
)   

// Transaction Friend List
type FirendList struct {
	
    EmailID string `json:"EmailID"`
	FrndEmailID string `json:"FrndEmailID"` 
	
    
} 

type FirendUpdateubscriptionList struct {
	
    Sender string `json:"Sender"`
	Text string `json:"Text"` 
	
    
} 
type FirendSubscriptionList struct {
	
		RequestorEmailID string `json:"RequestorEmailID"`
		TargetEmailID string `json:"TargetEmailID"` 	
		
	}

	type FirendBlockedSubscriptionList struct {
	
		RequestorEmailID string `json:"RequestorEmailID"`
		TargetEmailID string `json:"TargetEmailID"` 	
		
	}
// Custom response JSON object structure  
type FriendJsonObj struct {
	Success    string
	Friend   []string
	Count      int64  `json:"Count"`	 
	
}  	
type CreateFriendJsonObj struct {  //Structure for Json Object response
	Success bool
}

var Firends []FirendList  //Array for Friend 
var FirendSubscription []FirendSubscriptionList  //Array for Firend Subscription  
var FirendBlockedSubscription []FirendBlockedSubscriptionList  //Array for Friend Blocked Subscription
var FirendSubscriptionUpdate []FirendUpdateubscriptionList 
//Getting Common friends 
func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
			m[item] = true
	}

	for _, item := range b {
			if _, ok := m[item]; ok {
					c = append(c, item)
			}
	}
	return
}





func returnCommonFriendList(w http.ResponseWriter, r *http.Request){

	FriendArr:=[]string{}
	FriendArr_1:=[]string{}
	// Common_FriendArr:=[]string{}
	fmt.Println("Endpoint Hit: returnCommonFriendList")

	decoder := json.NewDecoder(r.Body)
			var data FirendList
			err := decoder.Decode(&data)
			if err != nil 	{
				panic(err)
			}
			fmt.Println(data.EmailID)

			for _, frned := range Firends {
				if frned.EmailID == data.EmailID {  //checking  email id is available or not in list		
					
					FriendArr = append(FriendArr,frned.FrndEmailID)  //Pushing into array		 
					
				}	
				
				if frned.EmailID == data.FrndEmailID {  //checking  email id is available or not in list		
					
					FriendArr_1 = append(FriendArr_1,frned.FrndEmailID)  //Pushing into array		 
					
				}
			}
			
			var result []string
			result=Intersection(FriendArr_1 , FriendArr)
			fmt.Println(FriendArr)
			fmt.Println(FriendArr_1)
			fmt.Println("Common Friends")
			 fmt.Println(result) 			//Calling Intersection function  for  getting common friends

			// Common_FriendArr = intersection(FriendArr, FriendArr_1)

			var FrndCommonCount int64=int64(len(result))

			FrndCommonJsonOBJ := FriendJsonObj{
				Success:    "True",
				Friend:result  ,
				Count: FrndCommonCount,
				
				
			}
			 
			var jsonData []byte
			jsonData, err1 := json.Marshal(FrndCommonJsonOBJ)
			if err1 != nil {  								//error handling
				log.Println(err1)
				log.Println(jsonData)
				json.NewEncoder(w).Encode(err1)
			}else{ 											//error is not occured 
				json.NewEncoder(w).Encode(FrndCommonJsonOBJ)		//Display JSON response in browser
			} 

			fmt.Println(FrndCommonJsonOBJ)


}

		// var Persons []PersonList 
		//return Friend List of specific person
		func returnPersonFriendList(w http.ResponseWriter, r *http.Request){
			// name:="Apple"
			FriendArr:=[]string{}
			
			
				
			fmt.Println("Endpoint Hit: returnPersonFriendList")
				decoder := json.NewDecoder(r.Body)
			var data FirendList
			err := decoder.Decode(&data)
			if err != nil 	{
				panic(err)
			}
			fmt.Println(data.EmailID)
			//  var s []string
			for _, frned := range Firends {
				if frned.EmailID == data.EmailID {  //checking  email id is available or not in list		
					
					FriendArr = append(FriendArr,frned.FrndEmailID)  //Pushing into array		 
					
				}		
			}
			var FrndCount int64=int64(len(FriendArr))

			FrndJsonOBJ := FriendJsonObj{
				Success:    "True",
				Friend:FriendArr  ,
				Count: FrndCount,
				
				
			}
			
			var jsonData []byte
			jsonData, err1 := json.Marshal(FrndJsonOBJ)
			if err1 != nil {  								//error handling
				log.Println(err1)
				log.Println(jsonData)
				json.NewEncoder(w).Encode(err1)
			}else{ 											//error is not occured 
				json.NewEncoder(w).Encode(FrndJsonOBJ)		//Display JSON response in browser
			} 
			
		}
		
		func returnUpdateFriendList(w http.ResponseWriter, r *http.Request){
			// name:="Apple"
			// FriendArr:=[]string{}
			
			 

// text := []byte(`Send me an email at foo@bar.com or foo@domain.fakesuffix.`)


				
			fmt.Println("Endpoint Hit: returnUpdateFriendList")
				decoder := json.NewDecoder(r.Body)
			var data FirendUpdateubscriptionList
			err := decoder.Decode(&data)
			if err != nil 	{
				panic(err)
			}
			if(data.Text!=""){
				fmt.Println(data.Text)
			} 

			text := []byte(data.Text)
validateHost := false

emails := emailaddress.Find(text, validateHost)
// fmt.Println(emails)
for _, e := range emails {
	//FirendSubscriptionList=append(FirendSubscriptionList,data.Sender,e)
	fmt.Println(e)
	
}
//json.NewEncoder(w).Encode(FirendSubscriptionList)

//
// 			// else{
			// 	fmt.Println(data.Text)
			// }
				
			
			// 	decoder1 := json.NewDecoder(r.Body)
			// var data1 FirendList
			// err2 := decoder1.Decode(&data1)
			// if err2 != nil 	{
			// 	panic(err2)
			// }
			 

			// fmt.Println(data1.EmailID)
			// //  var s []string
			// for _, frned := range FirendSubscription {
			// 	// if frned.RequestorEmailID == data.EmailID {  //checking  email id is available or not in list		
					
			// 		FriendArr = append(FriendArr,frned.TargetEmailID)  //Pushing into array		 
					
			// 	// }		
			// }
			// var FrndCount int64=int64(len(FriendArr))
			// json.NewEncoder(w).Encode(FirendSubscription)
			// FrndJsonOBJ := FriendJsonObj{
			// 	Success:    "True",
			// 	Friend:FriendArr  ,
			// 	Count: FrndCount,
				
				
			// }
			
			// var jsonData []byte
			// jsonData, err1 := json.Marshal(FrndJsonOBJ)
			// if err1 != nil {  								//error handling
			// 	log.Println(err1)
			// 	log.Println(jsonData)
			// 	json.NewEncoder(w).Encode(err1)
			// }else{ 											//error is not occured 
			// 	json.NewEncoder(w).Encode(FrndJsonOBJ)		//Display JSON response in browser
			// } 
			
		}

//Create New Friend
func create_NewFirend(w http.ResponseWriter, r *http.Request) {

	type CreateFriendJsonObj struct {  //Structure for Json Object response
		Success bool
	}

	
	var isSuccess bool=false
	fmt.Println("Endpoint Hit: create_NewFirend")	
    reqBody, _ := ioutil.ReadAll(r.Body)           // reading the request body
    var CreateNewFriend FirendList 
    json.Unmarshal(reqBody, &CreateNewFriend) 		//unmarshing body to JSON
    
	var isPersonValid bool =true 					//later we have to check person is valid or not	
	
	if(isPersonValid)	{
		Firends = append(Firends, CreateNewFriend) // Create new friend
	  json.NewEncoder(w).Encode(Firends)  
	 isSuccess=true 
	//fmt.Println(CreateNewFriend) 			// Display on Command prompt
			 

	}

	CreateFrndJsonOBJ := CreateFriendJsonObj{
		Success : isSuccess,			 
		}

		json.NewEncoder(w).Encode(CreateFrndJsonOBJ) // Display in Browser
		// json.NewEncoder(w).Encode(CreateNewFriend)
	 
// 	fmt.Println("Endpoint Hit: END")
   
} 

func Create_Subscription(w http.ResponseWriter, r *http.Request) {

	
	type CreateFriendJsonObj struct {  //Structure for Json Object response
		Success bool
	}
	
	var isSuccess bool=false
	fmt.Println("Endpoint Hit: Create_Subscription")

	reqBody, _ := ioutil.ReadAll(r.Body)           // reading the request body
    var CreateNewSubscription FirendSubscriptionList 
    json.Unmarshal(reqBody, &CreateNewSubscription) 		//unmarshing body to JSON
    
	var isPersonValid bool =true 					//later we have to check person is valid or not	
	
	if(isPersonValid)	{
		isSuccess=true 
		FirendSubscription = append(FirendSubscription, CreateNewSubscription) // Create new friend
	  json.NewEncoder(w).Encode(FirendSubscription) 
	fmt.Println(CreateNewSubscription) 			// Display on Command prompt
	
	}	
	CreateFrndJsonOBJ := CreateFriendJsonObj{
		Success : isSuccess,			 
		}

		json.NewEncoder(w).Encode(CreateFrndJsonOBJ) // Display in Browser 


}

func Blocked_Subscription(w http.ResponseWriter, r *http.Request) {

	
	type CreateFriendJsonObj struct {  //Structure for Json Object response
		Success bool
	}
	
	var isSuccess bool=false
	
	// var reuqestBody=r.Body
	reqBody, _ := ioutil.ReadAll(r.Body)           // reading the request body
	fmt.Println(reqBody)	
    var CreateBlockedSubscription FirendBlockedSubscriptionList 
    json.Unmarshal(reqBody, &CreateBlockedSubscription) 		//unmarshing body to JSON
    
	var isPersonValid bool =true 					//later we have to check person is valid or not	
	
	if(isPersonValid)	{
		isSuccess=true 
		FirendBlockedSubscription = append(FirendBlockedSubscription, CreateBlockedSubscription) // Create new friend
	//   json.NewEncoder(w).Encode(CreateBlockedSubscription) 
	//   json.NewEncoder(w).Encode(FirendBlockedSubscription) 
	fmt.Println(CreateBlockedSubscription) 			// Display on Command prompt
	
	}	
	CreateFrndJsonOBJ := CreateFriendJsonObj{
		Success : isSuccess,			 
		}

		json.NewEncoder(w).Encode(CreateFrndJsonOBJ) // Display in Browser 

		 
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		decoder := json.NewDecoder(r.Body)
		var data FirendBlockedSubscriptionList
		err := decoder.Decode(&data)
		if err != nil 	{
			panic(err)
		}
		fmt.Println(data.RequestorEmailID)

		for index, frnd := range FirendSubscription {
			if frnd.RequestorEmailID == data.RequestorEmailID && frnd.TargetEmailID == data.TargetEmailID{
				// updates our Articles array to remove the 
				// article
				FirendSubscription = append(FirendSubscription[:index], FirendSubscription[index+1:]...)
			}
		}
		fmt.Println(data.RequestorEmailID)
}



		func handleRequests() { 							//Request Section
			
			myRouter := mux.NewRouter().StrictSlash(true)    //Mapping Routes
			myRouter.HandleFunc("/CreateNewFriend", create_NewFirend).Methods("POST") 	//Calling create_NewFirend method
			myRouter.HandleFunc("/GetAllFriends", returnPersonFriendList).Methods("POST") //Calling returnPersonFriendList method
			myRouter.HandleFunc("/GetCommonFriend", returnCommonFriendList).Methods("POST") 	//Calling method for gettting common friends 
			myRouter.HandleFunc("/CreateUpdateSubscription", Create_Subscription).Methods("POST") 	//Calling method for gettting common friends 
			myRouter.HandleFunc("/BlockedUpdateSubscription", Blocked_Subscription).Methods("POST") 	//Calling method for gettting common friends 
			myRouter.HandleFunc("/UpdateSubscriptionList", returnUpdateFriendList).Methods("POST") 	//Calling method for gettting common friends 
			log.Fatal(http.ListenAndServe(":10000", myRouter)) //API Server rout 
		}



func main() { 								//Main Function
	
    fmt.Println("Rest API v2.0 - Mux Routers") //Just Display on Command propmt  for ensuring Api server is started
    Firends = []FirendList{
        FirendList{EmailID: "amol@gmail.com",FrndEmailID: "Frndamol1@gmail.com"},
		FirendList{EmailID: "amol@gmail.com",FrndEmailID: "Frndamol2@gmail.com"},
		FirendList{EmailID: "amol1@gmail.com",FrndEmailID: "Frndamol11@gmail.com"},
        FirendList{EmailID: "amol1@gmail.com",FrndEmailID: "Frndamol2@gmail.com"},
	}

	FirendSubscription = []FirendSubscriptionList{
        FirendSubscriptionList{RequestorEmailID: "amol@gmail.com",TargetEmailID: "Frndamol1@gmail.com"},
		FirendSubscriptionList{RequestorEmailID: "amol@gmail.com",TargetEmailID: "Frndamol2@gmail.com"},
		
	}

	FirendBlockedSubscription = []FirendBlockedSubscriptionList{
        FirendBlockedSubscriptionList{RequestorEmailID: "amol@gmail.com",TargetEmailID: "Frndamol1@gmail.com"},
		FirendBlockedSubscriptionList{RequestorEmailID: "amol@gmail.com",TargetEmailID: "Frndamol2@gmail.com"},
		
	}
	
// 	Persons = []PersonList{
//         PersonList{EmailID: "amol@gmail.com"},
//         PersonList{EmailID: "amol_1@gmail.com"},
//     }
    handleRequests() 						//Calling Handale Request method
}

