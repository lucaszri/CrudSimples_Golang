package handlers

import (
	"encoding/json"
	"net/http"
	"obra-crud/models"
	"obra-crud/repositories"
	"strconv"
	"strings"
)

type ItemCompraHandler struct {
	Repo *repositories.ItemCompraRepository
}

func NewItemCompraHandler(repo *repositories.ItemCompraRepository) *ItemCompraHandler {
	return &ItemCompraHandler{Repo: repo}
}

func (h *ItemCompraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(r.URL.Path, "/")

	id := strings.TrimPrefix(path, "/itens_compra/")
	hasID := id != "" && id != path

	switch {
	case path == "/itens_compra" && r.Method == http.MethodGet:
		h.GetAllItensCompra(w, r)
	case path == "/itens_compra" && r.Method == http.MethodPost:
		h.CreateItemCompra(w, r)
	case hasID && r.Method == http.MethodDelete:
		h.DeleteItemCompra(w, r, id)
	default:
		http.Error(w, "Rota ou método não encontrado", http.StatusNotFound)
	}
}

func (h *ItemCompraHandler) CreateItemCompra(w http.ResponseWriter, r *http.Request) {
	var itemCompra models.Item_compra
	err := json.NewDecoder(r.Body).Decode(&itemCompra)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.Repo.Create(itemCompra)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.FormatInt(id, 10)))
}

func (h *ItemCompraHandler) GetAllItensCompra(w http.ResponseWriter, r *http.Request) {
	itens, err := h.Repo.GetAll()	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(itens)
}

func (h *ItemCompraHandler) DeleteItemCompra(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	item, err := h.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}