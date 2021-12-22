package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(ctx *gin.Context, result interface{}) {
	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(ctx *gin.Context, response interface{}) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
