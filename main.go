package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type requestBody struct {
	Review string  `json:"review"`
	Score  float32 `json:"score"`
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func reviewScoreCalculator(reviews [][]string) <-chan requestBody {
	// Replace with your actual review score computation function or
	// an API call.
	ch := make(chan requestBody)
	go func() {
		defer close(ch)
		for idx, line := range reviews {
			if idx == 0 {
				continue
			}
			req := &requestBody{Review: line[1], Score: 7.5}
			ch <- *req
		}
	}()
	return ch
}

func chatCompletion(processedReview string, reviewScore float32) string {
	// The below code should work right out of the box, if you just change the
	// Request body to an actual prompt to the LLM
	// along with ollama being installed on your PC
	//req := &requestBody{Review: processedReview, Score: reviewScore}
	//reqBody, err := json.Marshal(req)
	//if err != nil {
	//	return ""
	//}
	//response, err := http.Post(
	//	"http://localhost:11434/api/generate",
	//	"application/json",
	//	bytes.NewBuffer(reqBody),
	//)
	//if err != nil {
	//	return ""
	//}
	//if response.StatusCode != http.StatusOK {
	//	return ""
	//}
	//defer response.Body.Close()
	//bytes, err := io.ReadAll(response.Body)
	//if err != nil {
	//	return ""
	//}
	//return string(bytes)
	//fmt.Printf("I have the review %s with a score %.2f\n", processedReview, reviewScore)
	sampleResponse := fmt.Sprintf(
		"For the review %s with score %.2f I think you're better off staying away from this product\n",
		processedReview,
		reviewScore,
	)
	return sampleResponse
}

func main() {
	csvs := readCsvFile("./test.csv")
	reviewScoreCh := reviewScoreCalculator(csvs)
	for msg := range reviewScoreCh {
		fmt.Println(chatCompletion(msg.Review, msg.Score))
	}
}
