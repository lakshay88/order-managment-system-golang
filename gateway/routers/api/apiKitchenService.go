package api

import "github.com/go-chi/chi"

type KitchenRoutes struct{}

func (k KitchenRoutes) Routes() chi.Router {
	r := chi.NewRouter()
	return r
}
