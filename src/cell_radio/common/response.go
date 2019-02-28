package common

type UsrpCfg struct {
	Config
	Watcher int
	ScanPsc string
	ScanTimeouts string
	ScanUlsc string
	Mode int
}

type AddManyResponse struct {
	UsrpCfg UsrpCfg
	Command int
}
