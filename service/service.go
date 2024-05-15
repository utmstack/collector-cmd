package serv

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"collector/utils"

	"github.com/kardianos/service"
)

type program struct {
	cmdRun  string
	cmdArgs []string
	path    string
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func (p *program) run() {
	err := utils.Execute(p.cmdRun, p.path, p.cmdArgs...)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
