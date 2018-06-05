package main

/*import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

//Book struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
	
}
//Author struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	
}
// Init Vooks var as a slice Book struct
var books []Book


//Create a New Book
func getBooks(w http.ResponseWriter , r  *http.Request){
	//json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter , r  *http.Request){

}
func createBook(w http.ResponseWriter , r  *http.Request){

}
func updateBook(w http.ResponseWriter , r  *http.Request){

}
func deleteBook(w http.ResponseWriter , r  *http.Request){

}

func main(){
	// Init Router
	r := mux.NetRouter()

	// Mock Data

	books = append(books , Book{ID :"1" , Isbn : "448743" , Title : "Book One" , Author: &Author{Firstname : "Sukum" , Lastname : "Nilphet"}})	
	books = append(books , Book{ID :"2" , Isbn : "448744" , Title : "Book Two" , Author: &Author{Firstname : "Sukum" , Lastname : "Nilphet"}})	

	// Router Handlers /enpoints
	r.HandlerFunc("/api/books",getBooks).Method("GET")
	r.HandlerFunc("/api/books/{id}",getBook).Method("GET")
	r.HandlerFunc("/api/books",createBook).Method("POST")
	r.HandlerFunc("/api/books/{id}",updateBook).Method("UPDATE")
	r.HandlerFunc("/api/books/{id}",deleteBook).Method("DELETE")
	log.Fatal(http.ListenAndServe(":8000" , r))

	
}*/
import("fmt"	
		"math")

func main(){
	fmt.Println("Test")
}