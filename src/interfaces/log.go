package interfaces

import (
	"LogGuardian/src/database"
	models "LogGuardian/src/models/log"
	models_database "LogGuardian/src/models/log/database"
	repository "LogGuardian/src/repository/sqllite"
	pacoteLog "log"
)

type InterfaceLog struct {
	log *models.Log
}

// NovaInterfaceDeLog: Cria uma instância da interface do Log
func NovaInterfaceDeLog(log *models.Log) *InterfaceLog {
	return &InterfaceLog{log}
}

// Criar: Salva o log no banco de dados
func (interfaceLog *InterfaceLog) Criar() {
	var log_sqlite models_database.Log_sqlite
	log_sqlite.ImportarLog(interfaceLog.log)

	db, err := database.Conectar()
	if err != nil {
		pacoteLog.Println("ocorreu um erro ao conectar com o banco de dados, error: ", err)
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioLog(db)

	id, err := repositorio.Criar(log_sqlite)
	if err != nil {
		pacoteLog.Println("ocorreu um erro ao criar o log, error:	", err)
	}
	if id == 0 {
		pacoteLog.Println("ocorreu algum erro ao criar o usuário, mas não retornou uma mensagem de erro. ")
	}
}
