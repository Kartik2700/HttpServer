package Controllers

import (
	"com.example/httpserver/DTO"
	"github.com/gin-gonic/gin"
	"fmt"
	 "net/http"
)

/*
	In easy terms, revisited:

 1. r.Body is the pipe through which the client sends you data.
 2. json.NewDecoder(r.Body) is like putting your ear to the pipe to listen.
 3. decoder.Decode(&userRequest) is like actually hearing and writing down the message.
    Once you hear it, it's passed you in the pipe.
 4. r.Body.Close() is like hanging up the phone after the message is delivered,
    making sure the pipe is properly closed and not left open consuming resources.
 5. The defer keyword just means: "No matter what happens in this function (success, error, panic),
    make sure to run r.Body.Close() right before the function completely finishes and exits."
    This guarantees cleanup.
*/
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	defer r.Body.Close()

// 	decoder := json.NewDecoder(r.Body)
// 	var userRequest DTO.UserRequest
// 	err := decoder.Decode(&userRequest)

// 	if err != nil {

// 		fmt.Printf("Error decoding request body: %v\n", err)

// 		userResponse := DTO.UserResponse{
// 			ID:      "",
// 			Message: "Invalid request body",
// 		}

// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(userResponse)
// 		return
// 	}

// 	userResponse := DTO.UserResponse{
// 		ID:      "12345",
// 		Message: "User created successfully",
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(userResponse)
// }


//using GIN
func CreateUser(c *gin.Context) {
	var userRequest DTO.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		fmt.Printf("Error decoding request body: %v\n", err)

		userResponse := DTO.UserResponse{
			ID:      "",
			Message: "Invalid request body",
		}

		c.JSON(http.StatusBadRequest, userResponse)

		return 
	}

	userResponse := DTO.UserResponse{
		ID:      "12345",
		Message: "User created successfully",
	}

	c.JSON(http.StatusCreated, userResponse)
}
