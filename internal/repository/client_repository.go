package repository

import (
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
	"net/http"
	"simple-crud/internal/model"
)

type ClientStorage struct {
	db *pgx.Conn
}

func NewClientStorage(db *pgx.Conn) *ClientStorage {
	return &ClientStorage{db: db}
}

func (c ClientStorage) Create(w http.ResponseWriter, r *http.Request) {
	var body model.Client
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sql, args, _ := sq.Insert("clients").
		Values(body.ID, body.Name, body.Surname, body.Lastname, body.Age, body.Weight, body.Height).
		ToSql()

	_, err = c.db.Exec(r.Context(), sql, args...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := model.CreateClientResponse{ID: body.ID}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, resp)
}

func (c ClientStorage) Read(w http.ResponseWriter, r *http.Request) {
	var body model.GetClientRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sql, args, _ := sq.Select("*").From("clients").Where(sq.Eq{"id": body.ID}).ToSql()

	var client model.Client
	err = pgxscan.Select(r.Context(), c.db, &client, sql, args...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, client)
}

func (c ClientStorage) Update(w http.ResponseWriter, r *http.Request) {
	var body model.Client
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sql, args, _ := sq.Update("clients").
		Set("id", body.ID).
		Set("name", body.Name).
		Set("surname", body.Surname).
		Set("lastname", body.Lastname).
		Set("age", body.Age).
		Set("height", body.Height).
		Set("weight", body.Weight).
		ToSql()

	_, err = c.db.Exec(r.Context(), sql, args...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, body)
}

func (c ClientStorage) Delete(w http.ResponseWriter, r *http.Request) {
	var body model.GetClientRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sql, args, _ := sq.Delete("*").From("clients").Where(sq.Eq{"id": body.ID}).ToSql()

	_, err = c.db.Exec(r.Context(), sql, args...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := model.DeleteClientResponse{ID: body.ID}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, resp)
}
