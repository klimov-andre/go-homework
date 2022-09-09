package handler

import "homework/internal/storage/facade"

type Handler struct {
	storage facade.StorageFacade
}

func NewHandler(storage facade.StorageFacade) *Handler {
	return &Handler{storage: storage}
}