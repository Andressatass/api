package controllers

import (
	banco "api/src/db"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuario no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(
			w,
			http.StatusUnprocessableEntity,
			err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(
			w,
			http.StatusBadRequest,
			err)
		return
	}

	if err = usuario.Preparar("cadastro"); err != nil {
		respostas.Erro(
			w,
			http.StatusBadRequest,
			err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(
			w,
			http.StatusInternalServerError,
			err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoDeUsuarios(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(
			w,
			http.StatusInternalServerError,
			err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	noeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepoDeUsuarios(db)
	usuarios, err := repositorios.Buscar(noeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuario no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuariosId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoDeUsuarios(db)
	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario atualiza um usuario no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuariosId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("edição"); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoDeUsuarios(db)
	if err = repositorio.Atualizar(usuarioID, usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario deleta um usuario no banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusUnauthorized, erro)
	// 	return
	// }

	// if usuarioID != usuarioIDNoToken {
	// 	respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu"))
	// 	return
	// }

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
