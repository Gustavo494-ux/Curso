// Teste de Unidade
package enderecos

import (
	"strings"
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	enderecosParaTeste := "Avenida Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecosParaTeste)

	if strings.ToLower(tipoDeEnderecoEsperado) != strings.ToLower(tipoDeEnderecoRecebido) {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s, e recebeu %s ",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}
}
