package profiling

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func Profiling() {
	fmt.Println("Inside main")
	// testing url for the profiling
	http.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {

		ch := make(chan int)
		go func() {
			obj := make(map[string]interface{})

			if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
				ch <- http.StatusBadRequest
				return
			}
			time.Sleep(time.Duration(rand.Intn(400)) * time.Microsecond)
			ch <- http.StatusAccepted
		}()
		for {
			select {
			case status := <-ch:
				w.WriteHeader(status)
			case <-time.After(time.Duration(200) * time.Microsecond):
				w.WriteHeader(http.StatusRequestTimeout)
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
