package handlers

import (
	"net/http"
	"obra-crud/repositories"
	"obra-crud/models"
	"encoding/json"
	"strconv"
)

type ProdutoHandler struct {
	Repo *repositories.ProdutoRepository
}

func NewProdutoHandler(repo *repositories.ProdutoRepository) *ProdutoHandler {
	return &ProdutoHandler{Repo: repo}
}

func (h *ProdutoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateProduto(w, r)
	case http.MethodGet:
		h.GetAllProdutos(w, r)
	case http.MethodPut:
		h.UpdateProduto(w, r)
	case http.MethodDelete:
		idStr := r.URL.Path[len("/produtos/"):]
		h.DeleteProduto(w, r, idStr)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func (h *ProdutoHandler) CreateProduto(w http.ResponseWriter, r *http.Request) {
	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 
	id, err := h.Repo.Create(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.FormatInt(id, 10)))
}

func (h *ProdutoHandler) GetAllProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, err := h.Repo.GetAll()	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func (h *ProdutoHandler) UpdateProduto(w http.ResponseWriter, r *http.Request) {
	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Repo.Update(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *ProdutoHandler) DeleteProduto(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}