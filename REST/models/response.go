package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response to reordener the code
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

//CreateDefaultResponse default
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

//Notfound default
func (this *Response) Notfound() { // puedes poner como parametro un VariableMessage sino es codigo duro como "messgae"
	this.Status = http.StatusNotFound
	this.Message = "Resource Not Found."
}
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.Notfound()
	response.Send()
}
func SendNotContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotContent()
	response.Send()
}
func (this *Response) NotContent() {
	this.Status = http.StatusNoContent
	this.Message = " Not content."
}

//Send enviar respuesta al cliente
func (this *Response) Send() {
	this.writer.Header().Set("Conten-type", this.contentType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output))
}
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func (this *Response) UnprocessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Message = "Unprocessable Not Found."
}

func (this *Response) SaveUser() {
	this.writer.Header().Set("Conten-type", this.contentType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output))
}
