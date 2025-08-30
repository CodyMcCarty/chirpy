package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	blWords := []string{"kerfuffle", "sharbert", "fornax"}
	const replacement = "****"
	cleanedBody := params.Body
	for _, wbad := range blWords {
		// build regex: case-insensitive whole word match
		re := regexp.MustCompile(`(?i)\b` + wbad + `\b`)
		cleanedBody = re.ReplaceAllString(cleanedBody, replacement)
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBody,
	})
}
