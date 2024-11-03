package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Alumnos struct {
	ID        int     `json:"id"`
	Nombres   string  `json:"nombres"`
	Apellidos string  `json:"apellidos"`
	Matricula string  `json:"matricula"`
	Promedio  float64 `json:"promedio"`
}

// Estructura para la respuesta
type Response struct {
	Alumnos []Alumnos `json:"alumnos"` // Contenedor para el slice de alumnos
}

var alumnos []Alumnos = []Alumnos{
	{ID: 1, Nombres: "Juan", Apellidos: "Pérez", Matricula: "2024-001", Promedio: 9.5},
	{ID: 2, Nombres: "María", Apellidos: "López", Matricula: "2024-002", Promedio: 8.7},
	{ID: 3, Nombres: "Pedro", Apellidos: "González", Matricula: "2024-003", Promedio: 7.8},
}
var response = Response{Alumnos: alumnos}

// Crear la respuesta con el slice de alumnos
// handler maneja las solicitudes HTTP
func getAlumnos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Serializar el slice de alumnos a JSON y escribirlo en la respuesta
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getAlumno(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	num, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "El parametro deber se un numero", http.StatusBadRequest)
	}
	alumno := alumnos[num-1]
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(alumno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func postAlumno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, has accedido a nuevo %s", r.URL.Path)
}

func putAlumno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, has accedido a actualizar %s", r.URL.Path)
}

func deleteAlumno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, has accedido a delete %s", r.URL.Path)
}

// main inicia el servidor HTTP
func main() {

	//http.HandleFunc("/", handler) // Asigna la función handler a la ruta "/"
	http.HandleFunc("GET /alumnos/", getAlumnos)
	http.HandleFunc("GET /alumnos/{id}", getAlumno)
	http.HandleFunc("POST /alumnos/", postAlumno)
	http.HandleFunc("PUT /alumnos/", putAlumno)
	http.HandleFunc("DELETE /alumnos/", deleteAlumno)
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Inicia el servidor en el puerto 8080
}
