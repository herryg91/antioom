package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/herryg91/antioom/src/antioom"
	"github.com/spf13/pflag"
	_ "github.com/tokopedia/dexter/profx/integration"
)

var memsize int
var commands []string

func initializeFlag() {
	pflag.IntVarP(&memsize, "memsize", "M", 500000, "Max Memory Size")
	pflag.StringArrayVarP(&commands, "commands", "C", []string{}, "Commands")
}

func startAntioom() {
	aoInstance := antioom.New(memsize, 1)
	for _, c := range commands {
		aoInstance.AddBashCommand(c)
	}
	aoInstance.Run()
}

func main() {
	initializeFlag()
	pflag.Parse()

	startAntioom()

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
				log.Println("[stop] Ctrl+C or Kill By SIGINT")
				exit_chan <- 0
			case syscall.SIGTERM: // kill -SIGTERM XXXX
				log.Println("[stop] Force Stop")
				exit_chan <- 0
			case syscall.SIGQUIT: // kill -SIGQUIT XXXX
				log.Println("[stop] Stop and Core Dump")
				exit_chan <- 0
			default:
				log.Println("[stop] Unknown Signal")
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}
