package handler

import (
	"encoding/json"
	"forum_backend/internal/project_error"
	"forum_backend/model"
	"io"
	"net/http"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.Context().Value("userID").(int)
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		h.Logger.Error("Couldn't read the body of a request in SignInHandler or body is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var comment model.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.services.IComment_service.InsertComment(comment, id); err != nil {
		switch err.(type) {
		case *project_error.UserError:
			userErr := err.(*project_error.UserError)
			http.Error(w, userErr.Error(), userErr.Status())
		case *project_error.InternalError:
			internalErr := err.(*project_error.InternalError)
			h.Logger.Error(internalErr.Error())
			http.Error(w, internalErr.Error(), http.StatusInternalServerError)
		default:
			serverErr := project_error.NewServerError("Internal Server Error")
			h.Logger.Error(serverErr.Error())
			http.Error(w, serverErr.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
