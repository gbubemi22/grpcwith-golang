// package middleware

// import (
// 	"net/http"
// 	"os"
// 
// 

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/writeconcern"

// 	"vendorService/internal/utils"
// )

// func ErrorHandlerMiddleware(c *gin.Context) {
// 	c.Next()

// 	serviceName := os.Getenv("SERVICE_NAME")

// 	// Check if there's an error to handle
// 	err := c.Errors.Last()
// 	if err != nil {
// 		// Default error response
// 		statusCode := http.StatusInternalServerError
// 		msg := "Internal Server Error"
// 		errorCode := "INTERNAL_SERVER_ERROR"

// 		// Determine the specific error type
// 		switch err.Type {
// 		case gin.ErrorTypeBind:
// 			statusCode = http.StatusBadRequest
// 			msg = "Invalid request data"
// 			errorCode = "BAD_REQUEST"
// 		case gin.ErrorTypeRender:
// 			statusCode = http.StatusInternalServerError
// 			msg = "Error rendering response"
// 			errorCode = "INTERNAL_SERVER_ERROR"
// 		default:
// 			// Handle custom errors
// 			if customErr, ok := err.Err.(*utils.CustomError); ok {
// 				statusCode = customErr.HTTPStatusCode
// 				msg = customErr.Message
// 				errorCode = "CUSTOM_ERROR"
// 			} else if mongoErr, ok := err.Err.(*mongo.CommandError); ok {
// 				switch mongoErr.Code {
// 				case 11000:
// 					// Handle duplicate key error
// 					statusCode = http.StatusConflict
// 					msg = "Duplicate value entered"
// 					if keyValueErr, ok := err.Err.(mongo.WriteException); ok {
// 						for _, we := range keyValueErr.WriteErrors {
// 							if we.Code == 11000 {
// 								msg = "Duplicate value entered for " + we.Message
// 							}
// 						}
// 					}
// 					errorCode = "DUPLICATE_ENTRY"
// 				// Add more cases for other MongoDB error codes
// 				default:
// 					msg = mongoErr.Message
// 					errorCode = "MONGO_ERROR"
// 				}
// 			} else if castErr, ok := err.Err.(*mongo.WriteConcernError); ok && castErr.Code == writeconcern.ErrCode {
// 				// Handle CastError equivalent
// 				statusCode = http.StatusNotFound
// 				msg = "No item found with id: " + castErr.Message
// 				errorCode = "NOT_FOUND"
// 			}
// 		}

// 		// Send JSON response
// 		c.JSON(statusCode, gin.H{
// 			"success":        false,
// 			"message":        msg,
// 			"httpStatusCode": statusCode,
// 			"error":          errorCode,
// 			"service":        serviceName,
// 		})
// 	}
// }
