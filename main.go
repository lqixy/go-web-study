package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })

	// http.ListenAndServe(":8000", nil)

	server := http.Server{
		Addr: "localhost:8000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, _ := ioutil.ReadFile("./index.html")
		w.Write(content)
	})

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// r.ParseMultipartForm(1024)

		// fileHeader := r.MultipartForm.File["upload"][0]
		// file, err := fileHeader.Open()

		file, _, err := r.FormFile("upload")

		if err == nil {
			data, err := ioutil.ReadAll(file)
			if err == nil {
				fmt.Fprintln(w, string(data))
			}
		}

		// r.ParseForm()
		// fmt.Fprintln(w, r.Form)
		// fmt.Fprintln(w, r.PostForm)

		// r.ParseMultipartForm(1024)
		// fmt.Fprintln(w, r.MultipartForm)

		fmt.Fprintln(w, r.FormValue("first-name"))

	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.RawQuery)
		fmt.Fprintln(w, r.URL.Query())
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	})

	server.ListenAndServe()
}

// import (
// 	"net/http"
// )

// type helloHandle struct{}

// func (m *helloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello"))
// }

// type aboutHandle struct{}

// func (h *aboutHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("about"))
// }

// func main() {

// 	h := helloHandle{}
// 	a := aboutHandle{}

// 	server := http.Server{
// 		Addr:    "localhostHandleFunc
// 		Handler: nil,
// 	}

// 	http.Handle("/hello", &h)
// 	http.Handle("/about", &a)

// 	server.ListenAndServe()
// 	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

// }

// // func index(w http.ResponseWriter, r *http.Request) {
// // 	content, _ := ioutil.ReadFile("./index.html")
// // 	w.Write(content)
// // }
