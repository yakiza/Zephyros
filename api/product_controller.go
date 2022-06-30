package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/product"
	"net/http"
)

var _ http.Handler = ProductController{}

type ProductController struct {
	chi.Router
	AddProductHandler product.AddProductHandler
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
	_ = h.AddProductHandler.AddProduct(kiteObj)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		_, err := w.Write([]byte("Duplicate found"))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func MakeProductController(addProductHandler product.AddProductHandler) ProductController {
	productController := ProductController{
		Router:            chi.NewRouter(),
		AddProductHandler: addProductHandler,
	}

	productController.Router.Post("/", productController.AddProduct)

	return productController
}
