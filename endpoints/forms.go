package endpoints

import (
	"analytics/dbgen"
	"encoding/json"
	"log"
	"net/http"
)

type QuestionInput struct {
	Question string `json:"question"`
	Type     string `json:"type"`
}

func (h Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var input QuestionInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	q_created, err := dbgen.New(h.Pool).CreateQuestion(r.Context(), dbgen.CreateQuestionParams{
		Question:   input.Question,
		AnswerType: input.Type,
	})

	response, err := json.Marshal(q_created)
	if err != nil {
		http.Error(w, "Error processing input", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
