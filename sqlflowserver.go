package main

type sqlFlowServer struct{}

func (s *sqlFlowServer) Run(*RunRequest, SQLFlow_RunServer) error {
	return nil
}
