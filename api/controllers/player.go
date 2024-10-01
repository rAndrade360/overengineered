package controllers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"terraform-mongodb-pratical-example/domain"
	"terraform-mongodb-pratical-example/services"
)

type PlayerController struct {
	Service *services.PlayerService
}

func (pc *PlayerController) Save(w http.ResponseWriter, r *http.Request) {
	var p domain.Player

	bodyb, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Bad Request"))
		w.WriteHeader(400)
		log.Println("Err: ", err)
		return
	}

	err = json.Unmarshal(bodyb, &p)
	if err != nil {
		w.Write([]byte("Bad Request"))
		w.WriteHeader(400)
		log.Println("Err: ", err)
		return
	}

	err = pc.Service.SavePlayer(context.Background(), &p)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(500)
		log.Println("Err: ", err)
		return
	}

	b, err := json.Marshal(p)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(500)
		log.Println("Err: ", err)
		return
	}

	w.Write(b)
}

func (pc *PlayerController) Get(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = 10
	}

	players, err := pc.Service.GetPlayers(context.Background(), page, perPage)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(500)
		log.Println("Err: ", err)
		return
	}

	b, err := json.Marshal(players)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(500)
		log.Println("Err: ", err)
		return
	}

	w.Write(b)
}
