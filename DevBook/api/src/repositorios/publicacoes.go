package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar cria uma nova publicação
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(" insert into publicacoes (titulo,conteudo,autor_id) values (?,?,?) ")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma publicacao
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		Select 
			p.*, 
			u.nick 
		from devbook.publicacoes p 
			inner join devbook.usuarios u on u.id = p.autor_id 
		where 
			p.id = ? `,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

// BuscarPublicacao traz as publicações que apareceriam no feed do usuário
func (repositorio Publicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		Select 
				p.*, 
				u.nick 
		from publicacoes p 
			inner join usuarios u on u.id = p.autor_id 
			INNER JOIN seguidores s ON s.usuario_id = u.id
		where 
			(u.id = ? or s.seguidor_id =?)  `,
		usuarioID,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// AtualizarPublicacao altera os dados de uma publicação
func (repositorio Publicacoes) AtualizarPublicacao(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		` update publicacoes set titulo =?, conteudo =? where id =? `,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// DeletarPublicacao exclui os dados de uma publicação
func (repositorio Publicacoes) DeletarPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		` delete from publicacoes where id =? `,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPublicacoesPorUsuario traz todas as publicações de um usuário especifico
func (repositorio Publicacoes) BuscarPublicacoesPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
        Select 
                p.*, 
                u.nick 
        from publicacoes p 
            inner join usuarios u on u.id = p.autor_id 
        where 
            u.id =? `,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// CurtirPublicacao adiciona 1 curtida a publicação curtida
func (repositorio Publicacoes) CurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		` UPDATE publicacoes SET curtidas = curtidas + 1 where id = ? `,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// DescurtirPublicacao remove 1 curtida da publicação curtida
func (repositorio Publicacoes) DescurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		` UPDATE publicacoes SET curtidas = 
		case 
			when curtidas > 0 then curtidas -1
			else 0
		end where id = ? `,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}