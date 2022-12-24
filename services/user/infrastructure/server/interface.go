package server

type Server interface {
	configHandler()
	Start()
}
