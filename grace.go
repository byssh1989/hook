package hook

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	// "time"
)

var (
	server   *http.Server
	listener net.Listener
	graceful = false
)

// GraceRun 软启动
func GraceRun(port string, handler http.Handler) {
	log.Println("grace run")

	server = &http.Server{Addr: port, Handler: handler}
	log.Printf("Listening and serving HTTP on %s\n", port)

	var err error
	if graceful {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		listener, err = net.Listen("tcp", server.Addr)
	}

	if err != nil {
		log.Errorf("GraceRun Err: %v \n", err)
	}

	go func() {
		err = server.Serve(listener)
		if err == http.ErrServerClosed {
			log.Printf("server closed \n")
			removePidFile()
			os.Exit(0)
		} else {
			log.Panicf("server.Serve err: %v \n", err)
		}
	}()

	listenSignal()
	log.Printf("server quit bye")
}

// ListenSignal 监听系统信号,并做出操作
func listenSignal() {
	ch := make(chan os.Signal, 1)
	listenList := []os.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGUSR2,
	}
	signal.Notify(ch, listenList...)

	var err error

	for {
		sig := <-ch
		log.Printf("signal: %v \n", sig)
		// ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
		ctx := context.Background()
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("start stop \n")
			signal.Stop(ch)
			err = server.Shutdown(ctx)
			if err != nil {
				log.Printf("shutdown err: %v \n", err)
			}
			log.Println("graceful shutdown")
			return
		case syscall.SIGUSR2:
			log.Println("reload")
			// 先更换listener, 然后再关闭老链接
			err = reload()
			if err != nil {
				log.Fatalf("grace ful restart err: %v", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload \n")
		}
	}
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)

	if !ok {
		return fmt.Errorf("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"start", "-g"}
	log.Printf("exec Command: %s %v", os.Args[0], args)
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}
