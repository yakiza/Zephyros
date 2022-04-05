package api

import (
	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros/databases/fakedb"
	"github.com/yakiza/Zephyros/kite"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	var db fakedb.DB
	addHandler := kite.NewAddHandler(db.Kite())

	r.Mount("/kite", MakeKiteController(addHandler))
	// r.Mount("/bar", MakeBarController())
	// r.Mount("/harness", MakeHarnessController())
	// r.Mount("/board", MakeBoardController())
	// r.Mount("wetsuit", MakeWetSuitController())

	return r
}
