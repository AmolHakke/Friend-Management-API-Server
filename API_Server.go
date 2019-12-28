//  package section
package main
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	 "github.com/gorilla/mux"
	 "io/ioutil"
	 //"bytes"
	 //"github.com/mcnijman/go-emailaddress"
	 //"gopkg.in/matryer/respond.v1"
	 	 
)   
// struct section
type FriendsRequest struct {
	Friends []string `json:"friends"`
}

type FriendListRequest struct {
	Email string `json:"email"`
}

// Transaction Friend List
type FriendList struct {
	
    EmailID string `json:"EmailID"`
	FriendEmailID string `json:"FriendEmailID"` 
	
    
}
type ReturnStatus struct {
	Success bool `json:"success"`
}

type ReturnFriendList struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type Subscription struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type Updates struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type UpdatesResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}


var friendsRequest FriendsRequest
var friendsListRequest FriendListRequest
var FriendsList []FriendList
var subscription Subscription
var updates Updates

//Sample homepage function to test that API server is working
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

// Create new friend connection
func createFriendsConnection(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new FriendRequest struct
    // append this to our FriendList array. 
	w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: FriendsConnection")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    var friend FriendList 
	friendsRequest = FriendsRequest{}
    json.Unmarshal(reqBody, &friendsRequest)
	 
	// Validations
    // Validate that email ids in proper format
    // Validate whether FriendRequest is not empty
    // Check whether they are friends already
     	
    // update our global FriendList array to include
    // our new friend connection
	friend.EmailID  = friendsRequest.Friends[0]
	friend.FriendEmailID = friendsRequest.Friends[1] 
    FriendsList = append(FriendsList, friend)
	
	
	returnStatus = ReturnStatus{true}
    json.NewEncoder(w).Encode(returnStatus)
	
}

// Get Friend List
func getFriendList(w http.ResponseWriter, r *http.Request) {
    // get the body of our GET request
    // unmarshal this into a new FriendListRequest struct
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: FriendsList")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    friendsListRequest = FriendListRequest{}
    json.Unmarshal(reqBody, &friendsListRequest)
	 
	// Validations
    // Validate that email ids in proper format
    // Validate whether FriendListRequest is not empty
        
	// Prepare some hardcoded response
	var returnFriendList ReturnFriendList
	returnFriendList.Success = true
	returnFriendList.Friends = append(returnFriendList.Friends,"abc@example")
	returnFriendList.Friends = append(returnFriendList.Friends,"cde@example")
	returnFriendList.Count = 2
	json.NewEncoder(w).Encode(returnFriendList)
	
}

// Get Common Friends
func getCommonFriends(w http.ResponseWriter, r *http.Request) {
    // get the body of our GET request
    // unmarshal this into a new FriendListRequest struct
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: CommonFriends")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    friendsRequest = FriendsRequest{}
    json.Unmarshal(reqBody, &friendsRequest)
	 
	// Validations
            
	// Prepare some hardcoded response
	var returnFriendList ReturnFriendList
	returnFriendList.Success = true
	returnFriendList.Friends = append(returnFriendList.Friends,"abc@example")
	returnFriendList.Friends = append(returnFriendList.Friends,"cde@example")
	returnFriendList.Count = 2
	json.NewEncoder(w).Encode(returnFriendList)
	
}

// Subscribe to Update
func subscribeUpdate(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Subscription struct
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: SubscribeUpdate")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    subscription = Subscription{}
    json.Unmarshal(reqBody, &subscription)
	 
	// Validations
    // Validate that email ids in proper format
    // Validate whether FriendRequest is not empty
    // Check whether they are friends already
     	
    returnStatus = ReturnStatus{true}
    json.NewEncoder(w).Encode(returnStatus)
	
}

// Block Updates
func blockUpdates(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new FriendRequest struct
    // append this to our FriendList array. 
	w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: BlockUpdate")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    subscription = Subscription{}
    json.Unmarshal(reqBody, &subscription)
	 
	// Validations
    returnStatus = ReturnStatus{true}
    json.NewEncoder(w).Encode(returnStatus)
	
}
// Receive Updates
func receiveUpdates(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Updates struct
    // append this to our FriendList array. 
	w.Header().Set("Content-Type", "application/json")
    fmt.Println("Endpoint Hit: ReceiveUpdate")
    var returnStatus ReturnStatus	
    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		returnStatus = ReturnStatus{false}
		json.NewEncoder(w).Encode(returnStatus)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
    updates = Updates{}
    json.Unmarshal(reqBody, &updates)
	 
	// Validations
    
    var updatesResponse UpdatesResponse
    updatesResponse.Success = true
	updatesResponse.Recipients = append(updatesResponse.Recipients,"abc@example")
	updatesResponse.Recipients = append(updatesResponse.Recipients,"cde@example") 	
    json.NewEncoder(w).Encode(updatesResponse)
	
}

// Existing code from above
func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/FriendsConnection", createFriendsConnection).Methods("POST")
	myRouter.HandleFunc("/SubscriptionUpdate", subscribeUpdate).Methods("POST")
	myRouter.HandleFunc("/FriendsList", getFriendList)
	myRouter.HandleFunc("/CommonFriends", getCommonFriends)
	myRouter.HandleFunc("/BlockUpdates", blockUpdates).Methods("POST")
	myRouter.HandleFunc("/ReceiveUpdates", receiveUpdates).Methods("POST")
	// finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func main() {
    fmt.Println("Friends API v2.0 - Powered by Mux Routers")
    handleRequests()
}
