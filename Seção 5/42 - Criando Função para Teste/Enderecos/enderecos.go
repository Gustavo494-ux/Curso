package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se um endereçoo tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	endereco = strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Inválido"
}
