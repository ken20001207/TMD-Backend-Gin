package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return errorHandlerT(gin.ErrorTypeAny)
}

func errorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err

			var parsedError *errorResponse

			switch err.Error() {
			case "EMAIL_OR_PASSWORD_WRONG":
				parsedError = &errorResponse{
					Code:    http.StatusNotFound,
					Message: "EMAIL_OR_PASSWORD_WRONG",
				}
			default:
				parsedError = &errorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}

			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()

			return
		}

	}
}
