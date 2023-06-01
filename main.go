package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		return
	}

	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	f.Flush()

	for i := 0; i < 100; i++ {
		chunk := fmt.Sprintf("id: %d\nevent: onProgress\ndata: {\"progressPercentage\": %d}\n\n", i, i*1)
		w.Write([]byte(chunk))
		f.Flush()
		time.Sleep(time.Millisecond * 100)
	}
	done := fmt.Sprintf("id: 10\nevent: done\ndata: {}\n\n")
	w.Write([]byte(done))
	f.Flush()
}

func main() {
	fs := http.FileServer(http.Dir("files"))
	http.HandleFunc("/hello", hello)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
		http.HandleFunc("/hello", hello)
		http.Handle("/files", http.FileServer(http.Dir("/Users/hwarren/Documents/vmware/projects/mess_around/server_side_events/files")))
		http.ListenAndServe(":8080", nil)
	*/
}
