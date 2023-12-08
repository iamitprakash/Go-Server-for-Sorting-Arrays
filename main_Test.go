func TestValidRequest(t *testing.T) {
	req := SortRequest{
		ToSort: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	body, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Error marshalling request: %v", err)
	}

	resp, err := http.Post("http://localhost:8000/process-single", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}


func TestInvalidRequestBody(t *testing.T) {
	body := []byte("invalid json")

	resp, err := http.Post("http://localhost:8000/process-single", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}


func TestEmptyRequestBody(t *testing.T) {
	body := []byte("")

	resp, err := http.Post("http://localhost:8000/process-single", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}


func TestNonExistentEndpoint(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/invalid-endpoint")
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}


func TestSortCorrectness(t *testing.T) {
	req := SortRequest{
		ToSort: [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
	}

	singleResp, err := http.Post("http://localhost:8000/process-single", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer singleResp.Body.Close()

	concurrentResp, err := http.Post("http://localhost:8000/process-concurrent", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
	defer concurrentResp.Body.Close()

	var singleResult, concurrentResult SortResponse
	json.NewDecoder(singleResp.Body).Decode(&singleResult)
	json.NewDecoder(concurrentResp.Body).Decode(&concurrentResult)

	for i, sortedArray := range singleResult.SortedArrays {
		if !reflect.DeepEqual(sortedArray, concurrentResult.SortedArrays[i]) {
			t.Errorf("Sorting results differ between single and concurrent: %v != %v", sortedArray, concurrentResult.SortedArrays[i])
		}
	}
}


func TestPerformanceComparison(t *testing.T) {
	req := SortRequest{
		ToSort: generateLargeArray(
