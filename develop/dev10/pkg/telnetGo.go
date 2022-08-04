package pkg

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type args struct {
	host    string
	port    string
	timeout int
}

type Telnet struct {
	args
}

func CreateTelnet(host, port string, timeout int) *Telnet {
	return &Telnet{
		args: args{
			host:    host,
			port:    port,
			timeout: timeout,
		},
	}
}

func request(conn net.Conn, osChan chan<- os.Signal, connChan chan<- error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		body, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osChan <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			connChan <- err
		}
		fmt.Fprintf(conn, body+"\n")
	}
}

func response(conn net.Conn, osChan chan<- os.Signal, connChan chan<- error) {
	for {
		reader := bufio.NewReader(conn)
		body, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osChan <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			connChan <- err
		}
		fmt.Print(body)
	}
}

func (t *Telnet) Run() error {
	address := net.JoinHostPort(t.args.host, t.args.port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(t.args.timeout)*time.Second)
	if err != nil {
		time.Sleep(time.Duration(t.args.timeout) * time.Second)
		return err
	}

	osChan := make(chan os.Signal, 1)
	connChan := make(chan error, 1)

	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)
	go request(conn, osChan, connChan)
	go response(conn, osChan, connChan)

	select {
	case <-osChan:
		conn.Close()
	case err = <-connChan:
		if err != nil {
			return err
		}
	}
	return nil
}
