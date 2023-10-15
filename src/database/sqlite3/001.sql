CREATE DATABASE LogGuardian
USE DATABASE LogGuardian

CREATE TABLE Log(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    codigoErro TEXT,
    tipo TEXT,
    nomePacote TEXT,
    nomeFuncao TEXT,
    LINHA INTEGER,
    mensagemErro TEXT,
    dadosAdicionais TEXT,
    dataHoraLog TEXT
)