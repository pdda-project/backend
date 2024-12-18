package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Konfigurasi stress test
	url := "http://localhost:8000/users"
	numRequests := 10000 // Total jumlah request
	numWorkers := 50     // Jumlah goroutine (workers)

	// Mulai stress test
	fmt.Println("Seeding...")

	start := time.Now()
	var wg sync.WaitGroup
	requestsPerWorker := numRequests / numWorkers

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for i := 0; i < requestsPerWorker; i++ {
				// Buat email unik untuk setiap request
				email := fmt.Sprintf("user_%s@example.com", uuid.NewString())
				requestBody := []byte(fmt.Sprintf(`{"email": "%s", "password": "secretpassword"}`, email))

				// Kirim HTTP POST request
				resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
				if err != nil {
					fmt.Printf("[Worker %d] Request failed: %v\n", workerID, err)
					continue
				}

				// Tampilkan hasil response
				fmt.Printf("[Worker %d] Received response: %s\n", workerID, resp.Status)
				resp.Body.Close()
			}
		}(w)
	}

	// Tunggu semua worker selesai
	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Stress test completed in %s\n", duration)
}
