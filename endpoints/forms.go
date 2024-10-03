package endpoints

import (
	"analytics/dbgen"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type QuestionInput struct {
	Question   string   `json:"question"`
	AnswerType string   `json:"answer_type"`
	Options    []string `json:"options"`
}

func (h Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var input QuestionInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	qCreated, err := dbgen.New(h.Pool).CreateQuestion(context.Background(), dbgen.CreateQuestionParams{
		Question:   input.Question,
		AnswerType: input.AnswerType,
		Options:    input.Options,
	})
	if err != nil {
		log.Println("Error inserting to the DB:", err)
		return
	}

	response, err := json.Marshal(qCreated)
	if err != nil {
		http.Error(w, "Error processing input", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
func (h Handler) CreateForm(w http.ResponseWriter, r *http.Request) {
	var input QuestionInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	qCreated, err := dbgen.New(h.Pool).CreateQuestion(context.Background(), dbgen.CreateQuestionParams{
		Question:   input.Question,
		AnswerType: input.AnswerType,
		Options:    input.Options,
	})
	if err != nil {
		log.Println("Error inserting to the DB:", err)
		return
	}

	response, err := json.Marshal(qCreated)
	if err != nil {
		http.Error(w, "Error processing input", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
