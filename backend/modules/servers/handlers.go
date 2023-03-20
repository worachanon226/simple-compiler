package servers

import "simple-compiler/backend/modules/controller"

func (s *Server) Maphandler() error {
	controller.ServerCheck(s.App)
	controller.Execute(s.App)
	controller.Compile(s.App)

	return nil
}
