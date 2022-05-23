package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/santidelosrios/platuit/ms_tuits/app/model"

	"github.com/go-chi/chi/v5"
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

// CreateTuit
func (h *Handler) CreateTuit(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, []byte("Error parsing the providing request"))
		return
	}

	tuitRequest := model.TuitRequest{}
	err = json.Unmarshal(data, &tuitRequest)

	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, []byte("Error parsing the providing request"))
		return
	}

	//TODO: Implement handler for POST /tuit
	if len(tuitRequest.Content) > 200 {
		writeJSONResponse(w, http.StatusBadRequest, []byte("Content of the tuit can not be greater than 200 characters"))
		return
	}

	query, err := h.db.Exec("INSERT INTO platuits (user_id, content) VALUES (?, ?)", tuitRequest.UserId, tuitRequest.Content)

	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, []byte("Could not create tuit"))
		return
	}

	id, _ := query.LastInsertId()

	writeJSONResponse(w, http.StatusCreated, []byte(strconv.Itoa(int(id))))
	return
}

//GetTuits
func (h *Handler) GetTuits(w http.ResponseWriter, r *http.Request) {
	//TODO: This should have limit and pagination
	tuits := model.GetPlatuitsResponse{}
	query, err := h.db.Query("SELECT * FROM platuits")

	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, []byte("Could not retrieve list of tuits"))
		return
	}

	defer query.Close()

	for query.Next() {
		tuit := model.Platuit{}

		query.Scan(&tuit.ID, &tuit.UserId, &tuit.Content)

		tuits = append(tuits, tuit)
	}

	logrus.Printf("Result of tuits %v", tuits)

	data, _ := json.Marshal(tuits)

	writeJSONResponse(w, http.StatusOK, data)
	return
}

func (h *Handler) GetTuit(w http.ResponseWriter, r *http.Request) {
	tuit := model.Platuit{}
	tuitID := chi.URLParam(r, "id")

	if tuitID == "" {
		writeJSONResponse(w, http.StatusBadRequest, []byte("id parameter is missing"))
		return
	}

	//TODO: Implement GET /tuit/{id}
	query, err := h.db.Query("SELECT * FROM platuits as p WHERE p.id = ?", tuitID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, []byte("Could not retrieve tuit"))
		return
	}

	defer query.Close()

	if query.Next() {
		query.Scan(&tuit.ID, &tuit.UserId, &tuit.Content)
	}

	data, _ := json.Marshal(tuit)

	writeJSONResponse(w, http.StatusOK, data)
	return

	//POST /interaction to register one visit
}

// writeJSONResponse Writes a response with a status code data
func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	w.Write(data)
}
