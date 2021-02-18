package main

import (
	"database/sql"
	"fmt"

	"github.com/fesyunoff/phone-book/pkg/book"
	"github.com/fesyunoff/phone-book/pkg/config"
	"github.com/fesyunoff/phone-book/pkg/storage/db"
	"github.com/fesyunoff/phone-book/pkg/transport"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	"net"
	"net/http"
	"os"
	"os/signal"
)

const (
	serviceName = "call-book"
)

func main() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	c := config.Config{
		HostDB:     "172.18.0.2",
		PortDB:     5432,
		UserDB:     "user",
		PasswordDB: "pass",
		NameDB:     "postgres",
		SchemaName: "phone",
		TableName:  "book",
	}
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	logger = kitlog.With(logger, "service", serviceName)
	logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestampUTC)
	logger = kitlog.With(logger, "caller", kitlog.Caller(5))
	connStatement := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.HostDB, c.PortDB, c.UserDB, c.PasswordDB, c.NameDB)
	connDB, err := sql.Open("postgres", connStatement)
	if err != nil {
		panic(err)
	}
	defer connDB.Close()
	db.PreparePostgresDB(connDB, &c)
	strg := db.NewPostgreBookStorage(connDB, &c)
	svc := book.NewService(strg)
	svcHandlers, err := transport.MakeHandlerJSONRPC(svc)
	bindAddr := fmt.Sprintf("%s:%d", "0.0.0.0", 8991)
	r := mux.NewRouter().StrictSlash(true)

	exitOnError(logger, err, "failed create handlers")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	r.PathPrefix("/").Handler(svcHandlers)

	ln, err := net.Listen("tcp", bindAddr)
	exitOnError(logger, err, "failed create listener")
	defer ln.Close()

	_ = level.Info(logger).Log("msg", "server listen on "+ln.Addr().String())

	go func() {
		_ = http.Serve(ln, r)
	}()

	<-sigint
}

func exitOnError(l kitlog.Logger, err error, msg string) {
	if err != nil {
		l.Log("err", err, "msg", msg)
		os.Exit(1)
	}
}
