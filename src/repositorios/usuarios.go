package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Usuarios representa um repo de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepoDeUsuarios cria um repositorio de usuarios
func NovoRepoDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usario no banco de dados
func (repo Usuarios) Criar(
	usuario models.Usuario) (uint64, error) {

	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Senha,
	)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuarios que atendem um filtro de
// nome ou nick.
func (repo Usuarios) Buscar(
	nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, err := repo.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuario de banco de dados
func (repo Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, err := repo.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nick,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuario no banco de dados
func (repo Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statment, err := repo.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.ID,
	); err != nil {
		return err
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}
