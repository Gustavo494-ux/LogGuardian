package GerenciadorDeArquivos_test

import (
	"LogGuardian/src/modules/GerenciadorDeArquivos"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	nomeArquivo   = "arquivoteste.txt"
	nomeDiretorio = "diretorioTeste"
)

func TestCriarArquivo(t *testing.T) {
	dir, _ := os.Getwd()
	err := GerenciadorDeArquivos.CriarArquivo(dir, nomeArquivo)
	if err != nil {
		t.Errorf("Falha ao criar o arquivo: %s", err)
	}
	os.Remove(nomeArquivo)
}

func TestAbrirArquivo(t *testing.T) {
	dir, _ := os.Getwd()
	conteudoEsperado := "Conteúdo do arquivo de teste"
	err := os.WriteFile(nomeArquivo, []byte(conteudoEsperado), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
	}
	defer os.Remove(nomeArquivo)

	conteudo, err := GerenciadorDeArquivos.AbrirArquivo(dir, nomeArquivo)
	if err != nil {
		t.Errorf("Falha ao abrir o arquivo: %s", err)
	}

	if conteudo != conteudoEsperado {
		t.Errorf("AbrirArquivo retornou conteúdo incorreto. Esperado: %s, Obtido: %s", conteudoEsperado, conteudo)
	}
}

func TestEscreverArquivo(t *testing.T) {
	dir, _ := os.Getwd()
	conteudo := "Conteúdo do arquivo de teste"
	err := GerenciadorDeArquivos.EscreverArquivo(dir, nomeArquivo, conteudo)
	if err != nil {
		t.Errorf("Falha ao escrever o arquivo: %s", err)
	}
	defer os.Remove(nomeArquivo)

	conteudoArquivo, err := os.ReadFile(nomeArquivo)
	if err != nil {
		t.Fatalf("Falha ao ler o arquivo de teste: %s", err)
	}

	if string(conteudoArquivo) != conteudo {
		t.Errorf("EscreverArquivo não escreveu o conteúdo esperado. Esperado: %s, Obtido: %s", conteudo, string(conteudoArquivo))
	}
}

func TestAnexarAoArquivo(t *testing.T) {
	conteudoInicial := "Conteúdo inicial"
	conteudoAnexado := "Conteúdo anexado"

	diretorioPath, _ := os.Getwd()
	caminhoCompleto := strings.ReplaceAll(filepath.Join(diretorioPath, nomeArquivo), "\\", "/")
	diretorioPath = diretorioPath + "\\"

	err := os.WriteFile(caminhoCompleto, []byte(conteudoInicial), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
	}
	defer os.Remove(caminhoCompleto)

	err = GerenciadorDeArquivos.AnexarAoArquivo(diretorioPath, nomeArquivo, conteudoAnexado)
	if err != nil {
		t.Errorf("Falha ao anexar ao arquivo: %s", err)
	}

	conteudoArquivo, err := os.ReadFile(caminhoCompleto)
	if err != nil {
		t.Fatalf("Falha ao ler o arquivo de teste: %s", err)
	}

	conteudoEsperado := conteudoInicial + conteudoAnexado
	if string(conteudoArquivo) != conteudoEsperado {
		t.Errorf("AnexarAoArquivo não anexou o conteúdo esperado. Esperado: %s, Obtido: %s", conteudoEsperado, string(conteudoArquivo))
	}
}

func TestExcluirArquivo(t *testing.T) {
	dir, _ := os.Getwd()
	caminhoCompleto := filepath.Join(dir, nomeArquivo)
	arquivo, err := os.Create(caminhoCompleto)
	if err != nil {
		t.Errorf("Falha ao criar o arquivo: %s", err)
	}
	arquivo.Close()
	err = GerenciadorDeArquivos.DeletarArquivo(dir, nomeArquivo)
	if err != nil {
		t.Errorf("Erro ao excluir o arquivo: %s", err)
	}

	_, err = os.Stat(nomeArquivo)
	if !os.IsNotExist(err) {
		t.Errorf("ExcluirArquivo não excluiu o arquivo como esperado")
	}
}

func TestRenomearArquivo(t *testing.T) {
	dir, _ := os.Getwd()
	novoNomeArquivo := "novoarquivo.txt"

	caminhoCompleto_inicial := strings.ReplaceAll(filepath.Join(dir, nomeArquivo), "\\", "/")
	caminhoCompleto_renomear := strings.ReplaceAll(filepath.Join(dir, novoNomeArquivo), "\\", "/")

	err := os.WriteFile(caminhoCompleto_inicial, []byte("Conteúdo do arquivo de teste"), 0644)
	if err != nil {
		t.Fatalf("Falha ao escrever o arquivo de teste: %s", err)
	}
	defer os.Remove(novoNomeArquivo)

	err = GerenciadorDeArquivos.RenomearArquivo(dir, nomeArquivo, novoNomeArquivo)
	if err != nil {
		t.Errorf("Falha ao renomear o arquivo: %s", err)
	}

	_, err = os.Stat(caminhoCompleto_inicial)
	if !os.IsNotExist(err) {
		t.Errorf("RenomearArquivo não renomeou o arquivo como esperado")
	}

	_, err = os.Stat(caminhoCompleto_renomear)
	if os.IsNotExist(err) {
		t.Errorf("RenomearArquivo não criou o arquivo renomeado como esperado")
	}
}

func TestObterListaDeArquivos(t *testing.T) {
	// Crie um diretório temporário para fins de teste
	diretorioAtual, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao buscar o caminho do diretório atual, err: %s", err)
	}

	caminhoCompleto := strings.ReplaceAll(filepath.Join(diretorioAtual, "test_directory"), "\\", "/")

	err = os.MkdirAll(caminhoCompleto, 0750)
	if err != nil {
		t.Fatalf("Falha ao criar o diretório temporário: %v", err)
	}
	defer os.RemoveAll("test_directory")

	// Crie arquivos de teste dentro do diretório
	arquivosTeste := []string{"arquivo1.txt", "arquivo2.txt", "arquivo3.txt"}
	for _, nomeArquivo := range arquivosTeste {
		caminhoArquivo := filepath.Join(caminhoCompleto, nomeArquivo)
		arquivo, err := os.Create(caminhoArquivo)
		if err != nil {
			t.Fatalf("Falha ao criar o arquivo de teste: %v", err)
		}
		defer arquivo.Close()
	}

	// Execute a função GetFileList no diretório de teste
	listaArquivos, err := GerenciadorDeArquivos.ObterListaDeArquivos(caminhoCompleto)
	if err != nil {
		t.Fatalf("Erro ao obter lista de arquivos: %v", err)
	}

	// Verifique se todos os arquivos de teste estão presentes na lista de arquivos retornada
	for _, nomeArquivo := range arquivosTeste {
		encontrado := false
		for _, arquivo := range listaArquivos {
			if arquivo == nomeArquivo {
				encontrado = true
				break
			}
		}
		if !encontrado {
			t.Errorf("Arquivo esperado ausente na lista de arquivos retornada: %s", nomeArquivo)
		}
	}

	// Verifique se não há arquivos extras na lista de arquivos retornada
	if len(listaArquivos) != len(arquivosTeste) {
		t.Errorf("Número de arquivos retornados é diferente do esperado. Esperado: %d, Retornado: %d", len(arquivosTeste), len(listaArquivos))
	}
}

func TestCriarDiretorio(t *testing.T) {
	// Especifique o caminho base para os novos diretórios
	caminhoBase := "./src/utility/teste"

	// Crie uma lista de nomes de pasta
	pastas := []string{"pasta1", "pasta2", "pasta3", "pasta4", "pasta5", "pasta6", "pasta7", "pasta8", "pasta9", "pasta10"}

	// Itere sobre a lista de pastas
	for _, pasta := range pastas {
		// Construa o caminho completo para cada diretório
		caminho := filepath.Join(caminhoBase, pasta)

		// Chame a função CriarDiretorio
		err := GerenciadorDeArquivos.CriarDiretorio(caminho)
		if err != nil {
			t.Errorf("Falha ao criar diretório: %v", err)
		}
	}

	caminhoBaseDividido := strings.Split(caminhoBase, "/")
	err := os.RemoveAll("./" + caminhoBaseDividido[1])
	if err != nil {
		t.Errorf("Falha ao remover o diretório: %v", err)
	}
}

func TestObterInformacoesDoArquivo(t *testing.T) {
	// Teste para um arquivo
	caminhoArquivo := "./src/utility/file.txt"
	infoArquivo, err := GerenciadorDeArquivos.ObterInformacoesDoArquivo(caminhoArquivo)
	if err != nil {
		t.Errorf("Falha ao obter informações do arquivo: %v", err)
	}

	if infoArquivo != nil {
		// Verifique se é um arquivo
		if !infoArquivo.Mode().IsRegular() {
			t.Errorf("Esperado um arquivo, obtido diretório")
		}
	}

	// Teste para um diretório
	caminhoDir := "./src/utility"
	infoDir, err := GerenciadorDeArquivos.ObterInformacoesDoArquivo(caminhoDir)
	if err != nil {
		t.Errorf("Falha ao obter informações do diretório: %v", err)
	}

	if infoDir != nil {
		// Verifique se é um diretório
		if !infoDir.Mode().IsDir() {
			t.Errorf("Esperado um diretório, obtido arquivo")
		}
	}
}
