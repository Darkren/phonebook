// Package server contains the JSONRPC server which uses
// websockets as a transport
package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	config "github.com/Darkren/go-config"
	"github.com/Darkren/graceful"
	"github.com/Darkren/phonebook/controllers"
	repoFabric "github.com/Darkren/phonebook/repositories/fabric/mysql"
	"github.com/Darkren/phonebook/wsjsonrpc"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Server is a JSONRPC WS server
type Server struct {
	userController  *controllers.UserController
	phoneController *controllers.PhoneController
	wsMaxMsgSize    int
	pongWaitSec     int
	wsUpgrader      websocket.Upgrader
}

// Start starts the server
func (s *Server) Start(config config.Config) {
	wsConfig, err := config.Section("ws")
	if err != nil {
		log.Fatalf("Got err reading config section: %v", err)
	}

	s.wsMaxMsgSize = wsConfig.MustGetInt("maxMsgSize")
	s.pongWaitSec = wsConfig.MustGetInt("pongWaitSec")

	s.wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  s.wsMaxMsgSize,
		WriteBufferSize: s.wsMaxMsgSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	dbConfigSection, err := config.Section("db")
	if err != nil {
		log.Fatalf("Error reading DB config: %v", err)
	}

	// get db connection
	db, err := s.setupDB(dbConfigSection)
	if err != nil {
		log.Fatalf("Couldn't establish connection to DB: %v", err)
	}

	repoFabric := repoFabric.New(db)

	userRepo := repoFabric.CreateUserRepository()
	phoneRepo := repoFabric.CreatePhoneRepository()

	s.userController = &controllers.UserController{UserRepo: userRepo}
	s.phoneController = &controllers.PhoneController{PhoneRepo: phoneRepo}

	port := config.MustGetInt("appPort")

	router := s.setupRoutes()

	server := http.Server{Addr: fmt.Sprintf(":%d", port), Handler: router}

	shutdown := graceful.Shutdown(&server)

	log.Printf("Server started listening on port: %v", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Couldn't start server on port %v: %v", port, err)
	}

	<-shutdown

	log.Println("Server stopped")
}

func (s *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("serveWs")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := s.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	s.handleWsConn(ws)
}

func wsping(ws *websocket.Conn, deadline time.Duration) error {
	return ws.WriteControl(websocket.PingMessage,
		[]byte{}, time.Now().Add(deadline*time.Second))
}

func wsclose(ws *websocket.Conn, deadline time.Duration) error {
	return ws.WriteControl(websocket.CloseMessage,
		[]byte{}, time.Now().Add(deadline*time.Second))
}

func (s *Server) handleWsConn(ws *websocket.Conn) {
	defer func() {
		deadline := 1 * time.Second
		wsclose(ws, deadline)
		time.Sleep(deadline)
		ws.Close()
	}()

	ws.SetReadLimit(int64(s.wsMaxMsgSize))
	ws.SetReadDeadline(time.Now().Add(time.Duration(s.pongWaitSec) * time.Second))

	go func() {
		ticker := time.Tick(time.Duration(s.pongWaitSec) * time.Second / 2)
		for range ticker {
			if err := wsping(ws, time.Duration(s.pongWaitSec)*time.Second); err != nil {
				log.Println("Ping failed:", err)
				break
			}
		}
		wsclose(ws, 1)
	}()

	rwc := &wsjsonrpc.ReadWriteCloser{WS: ws}
	rpcServer := rpc.NewServer()
	rpcServer.Register(s.userController)
	rpcServer.Register(s.phoneController)
	rpcServer.ServeCodec(jsonrpc.NewServerCodec(rwc))
}

func (s *Server) setupDB(dbConfig config.Config) (*sql.DB, error) {
	dbHost := dbConfig.MustGetString("host")
	dbPort := dbConfig.MustGetInt("port")
	dbUser := dbConfig.MustGetString("user")
	dbPassword := dbConfig.MustGetString("password")
	dbName := dbConfig.MustGetString("dbName")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// get db connection
	return sql.Open("mysql", connStr)
}

func (s *Server) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	sub := r.PathPrefix("/ws").Subrouter()

	sub.HandleFunc("", s.serveWs)
	sub.HandleFunc("/", s.serveWs)

	return s.setupStaticRoutes(r)
}

func (s *Server) setupStaticRoutes(r *mux.Router) *mux.Router {
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."+"/static/"))))
	r.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("."+"/static/"))))

	return r
}
