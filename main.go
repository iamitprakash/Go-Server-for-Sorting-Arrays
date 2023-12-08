package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"time"
)

type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64  `json:"time_ns"`
}

func main() {
	http.HandleFunc("/process-single", func(w http.ResponseWriter, r *http.Request) {
		var req SortRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		start := time.Now()
		sortedArrays := processSingle(req.ToSort)
		elapsed := time.Since(start)

		resp := SortResponse{
			SortedArrays: sortedArrays,
			TimeNs:       elapsed.Nanoseconds(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/process-concurrent", func(w http.ResponseWriter, r *http.Request) {
		var req SortRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		start := time.Now()
		sortedArrays := processConcurrent(req.ToSort)
		elapsed := time.Since(start)

		resp := SortResponse{
			SortedArrays: sortedArrays,
			TimeNs:       elapsed.Nanoseconds(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}

func processSingle(toSort [][]int) [][]int {
	var sortedArrays [][]int
	for _, arr := range toSort {
		sorted := make([]int, len(arr))
		copy(sorted, arr)
		sort.Ints(sorted)
		sortedArrays = append(sortedArrays, sorted)
	}
	return sortedArrays
}

func processConcurrent(toSort [][]int) [][]int {
	var wg sync.WaitGroup
	var sortedArrays [][]int
	ch := make(chan []int, len(toSort))

	for _, arr := range toSort {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()
			sorted := make([]int, len(arr))
			copy(sorted, arr)
			sort.Ints(sorted)
			ch <- sorted
		}(arr)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for sorted := range ch {
		sortedArrays = append(sortedArrays, sorted)
	}
	return sortedArrays
}
