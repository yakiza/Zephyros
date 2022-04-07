package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/kite"
)

var _ http.Handler = ProductController{}

type ProductController struct {
	chi.Router
	AddProductHandler kite.AddHandler
}

func (h *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
	}

	var kiteObj Zephyros.Product
	err := json.NewDecoder(r.Body).Decode(&kiteObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.AddProductHandler.Add(kiteObj)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *ProductController) GetProduct(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("title:%s", "What")))
	if err != nil {
		return
	}
}

func MakeProductController(addProductHandler kite.AddHandler) ProductController {
	productController := ProductController{
		Router:            chi.NewRouter(),
		AddProductHandler: addProductHandler,
	}

	productController.Router.Post("/", productController.AddProduct)
	productController.Router.Get("/", productController.GetProduct)

	return productController
}
