package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// sendErrorMessage sends an error message to the client
func SendErrorMessage(c *gin.Context, message string, code int) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{"error": message})
}

// sendSuccessMessage sends a success message to the client
func SendSuccessMessage(c *gin.Context, message string, data interface{}, code int) {
	// Set application/json as the content type
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{"message": message, "data": data})
}

// ToJSON converts a Go data structure to a JSON encoded byte slice
func ToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// FromJSON converts a JSON encoded byte slice to a Go data structure
func FromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// FromJSONRequest reads the request body and calls the FromJSON function
func FromJSONRequest(r *http.Request, v interface{}) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return FromJSON(data, v)
}

// ToJSONEncoder converts a Go data structure to a stream of JSON data
func ToJSONEncoder(w io.Writer, data interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

// FromJSONDecoder converts a stream of JSON data to a Go data structure
func FromJSONDecoder(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}

// Validate incoming JSON data using gin binding
func ValidateJSON(c *gin.Context, v interface{}) {
	if err := c.ShouldBindJSON(v); err != nil {
		SendErrorMessage(c, "Invalid JSON", 400)
		return
	}

}
