// Teste de Unidade
package enderecos

import (
	"strings"
	"testing"
)

type cenarioDeTeste struct {
	EnderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	CenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"RODOVIA dos Imigrantes", "Rodovia"},
		{"Praça Das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range CenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.EnderecoInserido)

		if strings.ToLower(tipoDeEnderecoRecebido) != strings.ToLower(cenario.retornoEsperado) {
			t.Errorf("O tipo recebido %s é diferente do tipo esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
