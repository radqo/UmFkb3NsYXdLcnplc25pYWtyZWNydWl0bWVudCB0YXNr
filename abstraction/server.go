package abstraction

// Server - server interface
type Server interface {
	Run(port string)
	Shutdown()
}
