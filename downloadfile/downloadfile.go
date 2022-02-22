package downloadfile


import (
	"log"
	"net/http"
)

func DownloadFile() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir("/Users/tiger/Downloads/"))))
}