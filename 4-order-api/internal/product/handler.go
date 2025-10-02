package product

import (
	"4-order-api/configs"
	"4-order-api/middleware"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}
type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.Handle("PATCH  /product/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /{id}", handler.GetById())
	router.HandleFunc("GET /products", handler.GetAll())

}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		fmt.Println(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product := NewProduct(body.Name, body.Description, body.Images)
		createProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createProduct, 201)
	}
}
func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
		if phone, ok := r.Context().Value(middleware.ContextPhonekey).(string); ok {
			fmt.Println(phone)
		}

		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id, err := handler.isId(w, r)
		if err != nil {
			return
		}
		_, err = handler.isExistById(w, id)
		if err != nil {
			return
		}
		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, product, 201)
	}
}
func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete")
		id, err := handler.isId(w, r)
		if err != nil {
			return
		}
		_, err = handler.isExistById(w, id)
		if err != nil {
			return
		}

		err = handler.ProductRepository.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GoTo")
		id, err := handler.isId(w, r)
		if err != nil {
			return
		}
		product, err := handler.isExistById(w, id)
		if err != nil {
			return
		}
		res.Json(w, product, 200)
	}
}
func (handler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetAll")
		pr, err := handler.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, pr, 200)
	}
}
func (handler *ProductHandler) isId(w http.ResponseWriter, r *http.Request) (uint, error) {
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}
	return uint(id), nil
}
func (handler *ProductHandler) isExistById(w http.ResponseWriter, id uint) (*Product, error) {
	product, err := handler.ProductRepository.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil, err
	}
	return product, nil

}
