# Go-Server-for-Sorting-Arrays

Go Server for Sorting Arrays

Objective:

Develop a Go server with two endpoints (/process-single and /process-concurrent) to demonstrate your skills in sequential and concurrent processing. The server should sort arrays provided in the request and return the time taken to execute the sorting in both sequential and concurrent manners.

Requirements:

Server Setup:
Create a Go server listening on port 8000
Implement two endpoints: /process-single for sequential processing and /process-concurrent for concurrent processing.

Input Format:
The server should accept a JSON payload with the following structure:
{
  "to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
}
Each element in the to_sort array is a sub-array that needs to be sorted.

Task Implementation:
For the /process-single endpoint, sort each sub-array sequentially.
For the /process-concurrent endpoint, sort each sub-array concurrently using Go's concurrency features (goroutines, channels).

Response Format:
Both endpoints should return a JSON response with the following structure:
{
  "sorted_arrays": [[sorted_sub_array1], [sorted_sub_array2], ...],
  "time_ns": "<time_taken_in_nanoseconds>"
}

Performance Measurement:
Measure the time taken to sort all sub-arrays in each endpoint in nanoseconds using Go's time package.

Dockerization:
Containerize your Go server using Docker.
Provide a Dockerfile for building the Docker image.

Submission:
Push the Docker image to Docker Hub.
Submit the source code repository link and the Docker Hub link for evaluation.

Evaluation Criteria:
Correctness of the sorting implementation in both sequential and concurrent methods.
Efficiency and performance comparison between the sequential and concurrent implementations.
Code quality, organization, and adherence to Go best practices.
Proper Dockerization of the application.





Running the server:

Build the Docker image:
docker build -t your-docker-username/server .
Push the Docker image to Docker Hub:
docker push your-docker-username/server
Run the Docker image:
docker run -p 8000:8000 your-docker-username/server
Testing the server:

Send a POST request to /process-single with the following JSON payload:
JSON
{"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}
Use code with caution. Learn more
Send a POST request to /process-concurrent with the same JSON payload.
Evaluation:

This code demonstrates the correct implementation of sequential and concurrent sorting with proper performance measurement. The code is well-organized and follows Go best
