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

func reviewScoreCalculator(review string, responseCh chan requestBody) {
	// Replace with your actual review score computation function or
	// an API call.
	req := &requestBody{Review: review, Score: 7.5}
	responseCh <- *req
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
	return "Based on the current review that I received, the product sucks! idk man I'm just an LLM."
}

func main() {
	reviewScoreCh := make(chan requestBody)
	csvs := readCsvFile("./test.csv")
	for idx, line := range csvs {
		if idx == 0 {
			continue
		}
		fmt.Println(idx)
		go reviewScoreCalculator(line[1], reviewScoreCh)
	}

	select {
	case msg := <-reviewScoreCh:
		llmResponse := chatCompletion(msg.Review, msg.Score)
		fmt.Println(llmResponse)
	}
}
