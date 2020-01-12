package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mnemonic111/poc_api_rest_go/controllers"
	"encoding/json"
	"github.com/mnemonic111/poc_api_rest_go/responses"
)

/* Funciones de Ejemplo */
func GetUsers(w http.ResponseWriter, r *http.Request) {
	controller := controllers.UserController{}
	usuario := controller.SetNewUser("Gonzalo","Mateos Glez-Sicilia", "gonzalomateosglezsicilia@gmail.com","blahblah");

	fmt.Println(usuario)

	//Generamos la respuesta.
	responseData := responses.GeneralResponse {}.GetInitResponse()
	tempData, _ := json.Marshal(usuario)
	responseData.SetDataResponse(string(tempData))

	//Enviamos la respuesta
	json.NewEncoder(w).Encode(responseData)
	//fmt.Fprintf(w, response)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo POST")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo PUT")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo DELETE")
}

/**  MAIN */
func main() {
	fmt.Println("Creando servidor")
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/v1/user", GetUsers).Methods("GET")       //El Get
	r.HandleFunc("/api/v1/user", PostUsers).Methods("POST")     //El Post
	r.HandleFunc("/api/v1/user", PutUsers).Methods("PUT")       //El Post
	r.HandleFunc("/api/v1/user", DeleteUsers).Methods("DELETE") //El Post

	//Creamos el servidor a partir de la libreria standar de Go http.
	servidor := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando en: ", servidor.Addr)
	servidor.ListenAndServe()

}
