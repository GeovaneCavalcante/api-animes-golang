package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GeovaneCavalcante/api-anime-golang/api/presenter"
	"github.com/GeovaneCavalcante/api-anime-golang/entity"

	"github.com/GeovaneCavalcante/api-anime-golang/usecase/anime"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getAnime(service anime.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading anime"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.GetAnime(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := presenter.Anime{
			ID:     data.ID,
			Banner: data.Banner,
			Name:   data.Name,
			Note:   data.Note,
			Genres: data.Genres,
			Year:   data.Year,
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

	})
}

func MakeBookHandlers(r *mux.Router, n negroni.Negroni, service anime.UseCase) {
	r.Handle("/v1/anime/{id}", n.With(
		negroni.Wrap(getAnime(service)),
	)).Methods("GET", "OPTIONS").Name("getAnime")
}
