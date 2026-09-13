package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/iconic-asf2910/student-api/internal/types"
	"github.com/iconic-asf2910/student-api/internal/utilis/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		slog.Info("creating a student")

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
