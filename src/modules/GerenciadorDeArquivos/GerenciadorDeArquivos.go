package GerenciadorDeArquivos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CriarArquivo : Cria um novo arquivo com o nome especificado
func CriarArquivo(caminhoDiretorio string, nomeArquivo string) error {
	arquivo, err := os.Create(strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeArquivo), "\\", "/"))
	if err != nil {
		return err
	}
	defer arquivo.Close()
	return nil
}

// AbrirArquivo : Abre um arquivo existente para leitura
func AbrirArquivo(caminhoDiretorio string, nomeArquivo string) (string, error) {
	dados, err := os.ReadFile(strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeArquivo), "\\", "/"))
	if err != nil {
		return "", err
	}
	return string(dados), nil
}

// EscreverArquivo : Escreve o conteúdo fornecido em um arquivo
func EscreverArquivo(caminhoDiretorio string, nomeArquivo string, conteudo string) error {
	caminhoCompleto := strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeArquivo), "\\", "/")
	err := os.WriteFile(caminhoCompleto, []byte(conteudo), 0600)
	if err != nil {
		return err
	}
	return nil
}

// AnexarAoArquivo : Anexa o conteúdo fornecido a um arquivo existente
func AnexarAoArquivo(caminhoDiretorio string, nomeArquivo string, conteudo string) error {
	arquivo, err := os.OpenFile(
		strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeArquivo), "\\", "/"),
		os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(conteudo)
	if err != nil {
		return err
	}
	return nil
}

// DeletarArquivo : Deleta um arquivo
func DeletarArquivo(caminhoDiretorio string, nomeArquivo string) error {
	caminhoCompleto := strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeArquivo), "\\", "/")
	err := os.Remove(caminhoCompleto)
	if err != nil {
		return err
	}
	return nil
}

// RenomearArquivo : Renomeia um arquivo
func RenomearArquivo(caminhoDiretorio string, nomeAntigo, nomeNovo string) error {
	caminhoCompletoAntigo := strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeAntigo), "\\", "/")
	caminhoCompletoNovo := strings.ReplaceAll(filepath.Join(caminhoDiretorio, nomeNovo), "\\", "/")
	err := os.Rename(caminhoCompletoAntigo, caminhoCompletoNovo)
	if err != nil {
		return err
	}
	return nil
}

// ObterListaDeArquivos : Retorna a lista de arquivos no diretório especificado
func ObterListaDeArquivos(diretorio string) ([]string, error) {
	listaArquivos := []string{}

	arquivos, err := os.ReadDir(diretorio)
	if err != nil {
		return nil, err
	}

	for _, arquivo := range arquivos {
		if !arquivo.IsDir() {
			listaArquivos = append(listaArquivos, arquivo.Name())
		}
	}

	return listaArquivos, nil
}

// CriarDiretorio : Cria um diretório no caminho especificado
func CriarDiretorio(caminho string) error {
	err := os.MkdirAll(caminho, 0750)
	if err != nil {
		return fmt.Errorf("erro ao criar o diretório: %v", err)
	}
	return nil
}

// ObterInformacoesDoArquivo : Retorna informações sobre um arquivo ou diretório especificado pelo caminho fornecido.
func ObterInformacoesDoArquivo(caminho string) (os.FileInfo, error) {
	// Converte o caminho em um caminho absoluto, se for relativo
	caminhoAbsoluto, err := filepath.Abs(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao resolver o caminho absoluto: %v", err)
	}

	// Verifica se o arquivo ou diretório existe
	_, err = os.Stat(caminhoAbsoluto)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Retorna nil quando o diretório não é encontrado
		}
		return nil, fmt.Errorf("erro ao obter informações do arquivo: %v", err)
	}

	// Recupera as informações do arquivo
	infoArquivo, err := os.Stat(caminhoAbsoluto)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter informações do arquivo: %v", err)
	}

	return infoArquivo, nil
}

// CriarDiretorioSeNaoExistir : Verifica se o diretório existe, caso contrário, ele será criado
func CriarDiretorioSeNaoExistir(caminho string) (err error) {
	infoDir, err := ObterInformacoesDoArquivo(caminho)
	if err != nil {
		err = fmt.Errorf("erro ao obter informações do diretório: %s", err)
	}

	if infoDir == nil {
		err = CriarDiretorio(caminho)
		if err != nil {
			err = fmt.Errorf("erro ao criar o diretório: %s", err)
		}
	}
	return
}

// CriarArquivoSeNaoExistir : Verifica se o arquivo existe, caso contrário, ele será criado
func CriarArquivoSeNaoExistir(caminho string) (err error) {
	diretorio, nomeArquivo := filepath.Split(caminho)
	infoArquivo, err := ObterInformacoesDoArquivo(caminho)
	if err != nil {
		err = fmt.Errorf("erro ao obter informações do arquivo: %s", err)
	}
	if infoArquivo == nil {
		err = CriarArquivo(diretorio, nomeArquivo)
		if err != nil {
			err = fmt.Errorf("erro ao criar o arquivo: %s", err)
		}
	}
	return
}

// ObterCaminhoDoDiretorio : recebe o caminho de um arquivo e extrai o caminho do diretório onde este arquivo será criado
func ObterCaminhoDoDiretorio(caminho string) string {
	caminhoDir := strings.Split(strings.ReplaceAll(caminho, "\\", "/"), "/")
	caminhoDir = append(caminhoDir[:len(caminhoDir)-1], caminhoDir[len(caminhoDir):]...)
	caminhoDirCriacao := ""
	for i, dir := range caminhoDir {
		if i > 0 {
			caminhoDirCriacao += "/"
		}
		caminhoDirCriacao += dir
	}
	return caminhoDirCriacao
}

// CriarDiretorioOuArquivoSeNaoExistir : Recebe o caminho de um arquivo e, se ele não existir, criará todos os diretórios e o arquivo em si.
func CriarDiretorioOuArquivoSeNaoExistir(caminho string) (err error) {
	err = CriarDiretorioSeNaoExistir(ObterCaminhoDoDiretorio(caminho))
	if err != nil {
		err = fmt.Errorf("erro CriarDiretorioSeNaoExistir : %s", err)
		return
	}

	err = CriarArquivoSeNaoExistir(caminho)
	if err != nil {
		err = fmt.Errorf("erro CriarArquivoSeNaoExistir : %s", err)
		return
	}
	return
}

// ObterCaminhoAtePasta : Retorna o caminho até o nível da pasta desejada em um caminho completo.
func ObterCaminhoAtePasta(caminho string, nomePasta string) (string, error) {
	// Procurar a posição da última ocorrência do nome da pasta no caminho
	indice := strings.LastIndex(strings.ToLower(caminho), nomePasta)
	if indice == -1 {
		return "", fmt.Errorf("a pasta '%s' não foi encontrada no caminho '%s'", nomePasta, caminho)
	}

	// Obter o caminho até a última ocorrência do nome da pasta
	caminhoAtePasta := caminho[:indice+len(nomePasta)]
	return caminhoAtePasta, nil
}

// ObterDiretorioFonte : Retorna o caminho absoluto para o diretório raiz do projeto
func ObterDiretorioFonte(DiretorioRaiz string) (caminhoDiretorioRaiz string, err error) {
	caminhoDiretorioAtual, err := os.Getwd()
	if err != nil {
		return
	}
	caminhoDiretorioRaiz, err = ObterCaminhoAtePasta(fmt.Sprintf("%s%s", caminhoDiretorioAtual, string(filepath.Separator)), DiretorioRaiz)
	return
}

// ObterCaminhoAbsolutoOuConcatenadoComRaizW : Função para verificar se o caminho fornecido é absoluto. caso contrário, será definido como o diretório raiz do projeto + o último diretório do caminho fornecido
func ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminho string, DiretorioRaiz string) (string, error) {
	caminho = filepath.FromSlash(caminho)
	if filepath.IsAbs(caminho) {
		return caminho, nil
	}

	DiretorioRaiz, err := ObterDiretorioFonte(DiretorioRaiz)
	if err != nil {
		return "", err
	}

	// Obter o último diretório e nome do arquivo do caminho
	diretorio, arquivo := filepath.Split(caminho)
	diretorio = strings.TrimSuffix(diretorio, string(filepath.Separator))
	arquivo = strings.TrimPrefix(arquivo, string(filepath.Separator))

	// Concatenar o último diretório e nome do arquivo com o caminho raiz
	caminhoAbsoluto := filepath.Join(DiretorioRaiz, diretorio, arquivo)
	return caminhoAbsoluto, nil
}
