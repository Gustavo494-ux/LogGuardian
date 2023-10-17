package interfaces

import (
	"LogGuardian/src/database"
	enum "LogGuardian/src/enum/log"
	models "LogGuardian/src/models/log"
	models_database "LogGuardian/src/models/log/database"
	repository "LogGuardian/src/repository/sqllite"
	pacoteLog "log"
)

type InterfaceLog struct {
	Log *models.Log
}

// NovaInterfaceDeLog: Cria uma instância da interface do Log
func NovaInterfaceDeLog(log *models.Log) *InterfaceLog {
	return &InterfaceLog{log}
}

// BuscarPorTipo: Busca todos os logs do tipo informado
func (interfaceLog *InterfaceLog) Criar() {
	var log_sqlite models_database.Log_sqlite
	log_sqlite.ImportarLog(*interfaceLog.Log)

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

// BuscarPorTipo: Busca todos os logs do tipo informado
func (InterfaceLog *InterfaceLog) BuscarPorTipo(tipoLog string) (log []models.Log) {
	db, err := database.Conectar()
	if err != nil {
		pacoteLog.Println("ocorreu um erro ao conectar com o banco de dados, error: ", err)
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioLog(db)
	log_sqlite, err := repositorio.BuscarPorTipo(enum.TipoLog(tipoLog))
	if err != nil {
		pacoteLog.Println("ocorreu um erro ao buscar os logs pelo tipo, error: ", err)
	}

	log = make([]models.Log, len(log_sqlite))
	for i := 0; i < len(log_sqlite); i++ {
		log[i] = log_sqlite[i].ExportarLog()
	}
	return
}
