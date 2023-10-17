
CREATE TABLE Log(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    codigoErro TEXT,
    tipo TEXT,
    nomePacote TEXT,
    nomeFuncao TEXT,
    linha INTEGER default 0,
    mensagemRetorno Text not null,
    mensagemErro TEXT not null,
    dadosAdicionais TEXT,
    dataHoraLog TEXT not null
)