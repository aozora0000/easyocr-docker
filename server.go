package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(404)
		w.Write([]byte("Support Only POST Method"))
		return
	case "POST":
		uploadFile(w, r)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("File Size Over")
		fmt.Println(err)
		return
	}
	// Get handler for filename, size and headers
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	dest := path.Join("/tmp/", strconv.FormatInt(time.Now().Unix(), 10))
	f, err := os.Create(dest)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Create TmpFile Failed")
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	cmd := exec.Command("python", "ocr.py", dest)
	result, err := cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Command Execute Failed"))
		fmt.Printf("Command: %s", cmd.String())
		fmt.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(result)
}

func main() {
	// Upload route
	http.HandleFunc("/", uploadHandler)

	//Listen on port 8080
	fmt.Println("Serve Start :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
