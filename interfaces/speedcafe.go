package interfaces

type ISpeedCafe interface {
	Init() error
	InitConfig() error
	Run()
	SetConfigPath(path string)
	SetPort(port string)
}
