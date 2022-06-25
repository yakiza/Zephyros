package api

import (
	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros/product"
	"github.com/yakiza/Zephyros/storage/fakedb"
	"go.uber.org/zap"
	"net/http"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	var db fakedb.DB
	logger, _ := zap.NewProduction()
	addHandler := product.NewAddProductHandler(logger, db.Product())

	r.Mount(HealthCheckMountPoint, HealthCheckController())
	r.Mount(ProductMountPoint, MakeProductController(addHandler))

	return r
}
