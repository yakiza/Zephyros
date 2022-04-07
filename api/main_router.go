package api

import (
	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros/databases/fakedb"
	"github.com/yakiza/Zephyros/kite"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	var db fakedb.DB
	logger, _ := zap.NewProduction()
	addHandler := kite.NewAddHandler(logger, db.Kite())

	r.Mount("/kite", MakeProductController(addHandler))
	// r.Mount("/bar", MakeBarController())
	// r.Mount("/harness", MakeHarnessController())
	// r.Mount("/board", MakeBoardController())
	// r.Mount("wetsuit", MakeWetSuitController())

	return r
}
