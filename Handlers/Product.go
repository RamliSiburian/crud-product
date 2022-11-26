package Handlers

import (
	"encoding/json"
	"net/http"
	productDto "nutecth/Dto/Product"
	Dto "nutecth/Dto/Result"
	"nutecth/Models"
	"nutecth/Repositories"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerproduct struct {
	ProductRepository Repositories.ProductRepository
}

func HandlerProduct(ProductRepository Repositories.ProductRepository) *handlerproduct {
	return &handlerproduct{ProductRepository}
}

func (h *handlerproduct) FindProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range products {
		products[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerproduct) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product Models.Product
	product, err := h.ProductRepository.GetProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerproduct) GetProductByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	var product []Models.Product
	product, err := h.ProductRepository.GetProductByUser(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range product {
		product[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerproduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataUpload := r.Context().Value("dataFile")
	filepath := dataUpload.(string)

	harga_beli, _ := strconv.Atoi(r.FormValue("harga_beli"))
	harga_jual, _ := strconv.Atoi(r.FormValue("harga_jual"))
	stok, _ := strconv.Atoi(r.FormValue("stok"))

	request := productDto.ProductRequest{
		Nama:      r.FormValue("nama"),
		HargaBeli: harga_beli,
		HargaJual: harga_jual,
		Stok:      stok,
		UserID:    userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "halloCorona/articleImage"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	product := Models.Product{
		Image:     filepath,
		Nama:      request.Nama,
		HargaBeli: request.HargaBeli,
		HargaJual: request.HargaJual,
		Stok:      request.Stok,
		UserID:    userId,
	}

	product, err = h.ProductRepository.CreateProduct(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProductById(product.ID)

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerproduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteProduct, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: deleteProduct}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerproduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filepath := ""
	productImage := r.Context().Value("dataFile")
	if productImage != nil {
		filepath = productImage.(string)
	}

	harga_beli, _ := strconv.Atoi(r.FormValue("harga_beli"))
	harga_jual, _ := strconv.Atoi(r.FormValue("harga_jual"))
	stok, _ := strconv.Atoi(r.FormValue("stok"))

	request := productDto.UpdateProductRequest{
		Nama:      r.FormValue("nama"),
		HargaBeli: harga_beli,
		HargaJual: harga_jual,
		Stok:      stok,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProductById(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "halloCorona/userImage"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	if request.Nama != "" {
		product.Nama = request.Nama
	}
	if request.HargaBeli != 0 {
		product.HargaBeli = request.HargaBeli
	}
	if request.HargaJual != 0 {
		product.HargaJual = request.HargaJual
	}
	if request.Stok != 0 {
		product.Stok = request.Stok
	}
	if filepath != "" {
		product.Image = filepath
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
