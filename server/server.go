package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"path"
	"path/filepath"
	"time"

	client "cpfi.client"
)

type Server struct {
	coordinator *Coordinator
}

func NewServer() *Server {
	listener := &Server{
		coordinator: NewCoordinator(),
	}
	listener.Start()
	return listener
}

func (s *Server) Start() {
	logger.Infof("Start Server....")
}

func (s *Server) Hello(request *client.CPFI_request, response *client.CPFI_response) error {
	fmt.Println(request.Text)
	return nil
}

func (s *Server) StartWorking(request *client.CPFI_request, response *client.CPFI_response) error {
	blockingCh := make(chan string)
	notification := &StartWorkingNotification{
		blockingCh: blockingCh,
		configPath: request.Text,
	}
	s.coordinator.raftNotificationCh <- notification
	<-blockingCh
	*response = client.CPFI_response{Ok: true}
	return nil
}

func (s *Server) AppendEntries(request *client.CPFI_append_request, response *client.CPFI_append_response) error {
	msg := &request.Msg
	ok := true
	blockingCh := make(chan string)
	notification := &AppendEntriesNotification{
		msg:        msg,
		blockingCh: blockingCh,
		who:        request.Who,
		when:       request.When,
	}
	s.coordinator.raftNotificationCh <- notification
	if <-blockingCh == Omit {
		ok = false
	}
	*response = client.CPFI_append_response{Ok: ok, Msg: *msg}
	return nil
}

// global
var logger *Logger
var mode string
var targetSystem string
var projectPath string
var interval int

type Config struct {
	Mode     string `json:"mode"`
	System   string `json:"system"`
	Interval int    `json:"interval"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	projectPath = path.Dir(path.Dir(exePath))
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("无法获取当前工作目录:", err)
	}
	// 获取当前日期和时间
	currentTime := time.Now().Format("2006-01-02_15:04")
	// 构建日志文件名
	logFileName := "conan_" + currentTime + ".log"
	logFilePath := filepath.Join(currentDir, "../logs", logFileName)
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal("Unable to create log file:", err)
	}
	defer logFile.Close()

	logger = NewLogger("", log.Ldate|log.Ltime, Info, logFile)

	args := os.Args
	var configPath string
	if len(args) < 2 {
		logger.Errorf("Usage: ./conan-server config_path")
		return
	} else {
		configPath = args[1]
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// 解析 JSON 数据
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	mode = config.Mode
	targetSystem = config.System
	interval = config.Interval
	logger.Infof("Start in %s mode\n", mode)
	logger.Infof("Target System is %s\n", targetSystem)

	rpc.Register(NewServer())
	log.Println("setting up connection...")
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	checkError(err)
	inbound, err := net.ListenTCP("tcp", addr)
	checkError(err)
	rpc.Accept(inbound)
}
