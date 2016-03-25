package interfaces

type ISpeedCafe interface {
	Init() error
	Run()
	SetConfigPath(path string)
	SetPort(port string)
}
