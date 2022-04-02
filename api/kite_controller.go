package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/kite"
)

var _ http.Handler = KiteController{}

type KiteController struct {
	chi.Router
	AddKiteHandler kite.AddHandler
}

func (h *KiteController) AddKite(_ http.ResponseWriter, r *http.Request) {
	var kiteObj Zephyros.Kite
	err := json.NewDecoder(r.Body).Decode(&kiteObj)
	if err != nil {
		return
	}
	err = h.AddKiteHandler.Add(kiteObj)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *KiteController) GetKite(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("title:%s", "What")))
	if err != nil {
		return
	}
}

func MakeKiteController(addKiteHandler kite.AddHandler) KiteController {
	kiteController := KiteController{
		Router:         chi.NewRouter(),
		AddKiteHandler: addKiteHandler,
	}

	kiteController.Router.Post("/", kiteController.AddKite)
	kiteController.Router.Get("/", kiteController.GetKite)

	return kiteController
}
