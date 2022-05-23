package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/santidelosrios/platuit/ms_reaction/app/model"

	"github.com/sirupsen/logrus"
)

//Handler struct for endpoints handlers
type Handler struct {
	client *http.Client
	db     *sql.DB
}

//NewHandler - Creates a new handler
func NewHandler(client *http.Client, db *sql.DB) *Handler {
	return &Handler{client: client, db: db}
}

// SendReaction
func (h *Handler) CreateVisitToTuit(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, []byte("Error parsing the providing request"))
		return
	}

	reactionRequest := model.ReactionVisitRequest{}
	err = json.Unmarshal(data, &reactionRequest)

	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, []byte("Error parsing the providing request"))
		return
	}

	query, err := h.db.Query("SELECT id, visits FROM platuit_reaction as pr WHERE pr.user_id = ? AND pr.tuit_id = ?", reactionRequest.UserId, reactionRequest.TuitId)

	if err != nil {
		logrus.Printf("Error in retrieving %s", err.Error())
		writeJSONResponse(w, http.StatusInternalServerError, []byte("Error retrieving reactions"))
		return
	}

	defer query.Close()

	if query.Next() {
		_, err := h.db.Exec("UPDATE platuit_reaction as pr SET visits = visits + 1 WHERE pr.user_id = ? AND pr.tuit_id = ?", reactionRequest.UserId, reactionRequest.TuitId)
		if err != nil {
			logrus.Printf("Error in updating %s", err.Error())
			writeJSONResponse(w, http.StatusInternalServerError, []byte("Error updating tuit reaction "+err.Error()))
			return
		}
	} else {
		_, err := h.db.Exec("INSERT INTO platuit_reaction (user_id, tuit_id, visits, like, dislike) VALUES (?, ?, 1, 0, 0)", reactionRequest.UserId, reactionRequest.TuitId)
		if err != nil {
			logrus.Printf("Error in creating %s", err.Error())
			writeJSONResponse(w, http.StatusInternalServerError, []byte("Error creating reactions "+err.Error()))
			return
		}
	}

	writeJSONResponse(w, http.StatusOK, []byte(""))
	return
}

// writeJSONResponse Writes a response with a status code data
func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	w.Write(data)
}
