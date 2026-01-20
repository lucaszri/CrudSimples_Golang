package handlers

import (
	"encoding/json"
	"net/http"
	"obra-crud/models"
	"obra-crud/repositories"
	"strconv"
	"strings"
)

type ProdutoHandler struct {
	Repo *repositories.ProdutoRepository
}

func NewProdutoHandler(repo *repositories.ProdutoRepository) *ProdutoHandler {
	return &ProdutoHandler{Repo: repo}
}

func (h *ProdutoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(r.URL.Path, "/")
	
	id := strings.TrimPrefix(path, "/produtos/")
	hasID := id != "" && id != path
	
	// Combinar path + método em um switch
	switch {
	case path == "/produtos" && r.Method == http.MethodGet:
		h.GetAllProdutos(w, r)
	case path == "/produtos" && r.Method == http.MethodPost:
		h.CreateProduto(w, r)
	case hasID && r.Method == http.MethodGet:
		h.GetProdutoByID(w, r, id)
	case hasID && r.Method == http.MethodPut:
		h.UpdateProduto(w, r, id)
	case hasID && r.Method == http.MethodDelete:
		h.DeleteProduto(w, r, id)
	default:
		http.Error(w, "Rota ou método não encontrado", http.StatusNotFound)
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

func (h *ProdutoHandler) GetProdutoByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	produto, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func (h *ProdutoHandler) UpdateProduto(w http.ResponseWriter, r *http.Request, idStr string) {
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