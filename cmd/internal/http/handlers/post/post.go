package post

import (
	"encoding/json"
	"errors"
	"github/SpaceBuckett/books-api/cmd/internal/types"
	"github/SpaceBuckett/books-api/cmd/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
)

/// WILL BE DOING CRUD HERE
// TODO:
// TODO: 1. CREATE NEW POST

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("CREATING A NEW TRASH!")

		var post types.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return

		}
		w.Write([]byte("CREATING A NEW POST..."))
		// response.WriteJson(w, http.StatusCreated)
	}
}

/// NOTES:
/// FOR EMPTY BODY TO DETECT
/// USE THE io.EOF FROM THE GO's ERRORS PACKAGE
