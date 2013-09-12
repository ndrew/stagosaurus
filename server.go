package stagosaurus

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func previewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
	// defer r.Body.Close()
}

func stopServer(w http.ResponseWriter, req *http.Request) {
	responseString := "Bye-bye"

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseString)))
	io.WriteString(w, responseString)

	f, canFlush := w.(http.Flusher)
	if canFlush {
		f.Flush()
	}

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		//fmt.Printf("error while shutting down: %v", err)
	}

	conn.Close()

	println("Shutting down")
	os.Exit(0)
}

func runServer(dir string, port string) { // "."
	//port.star

	http.HandleFunc("/exit", stopServer)
	http.HandleFunc("/preview", previewHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}
