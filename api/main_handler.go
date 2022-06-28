package api

import (
	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros/databases/fakedb"
	"github.com/yakiza/Zephyros/product"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	var db fakedb.DB
	logger, _ := zap.NewProduction()
	addHandler := product.NewAddHandler(logger, db.Product())

	r.Mount(ProductMountPoint, MakeProductController(addHandler))

	return r
}
