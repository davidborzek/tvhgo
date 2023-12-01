package e2e_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/cmd/server"
	"github.com/urfave/cli/v2"
)

func generateSqliteInMemoryDsn() string {
	return fmt.Sprintf("file:tvhgo_e2e_%d?mode=memory&cache=shared", time.Now().Unix())
}

func addTestData(dsn string) {
	dbConn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	runSqlScript(dbConn, "./sql/init.sql")
}

func waitForServiceIsReady() {
	c := http.Client{Timeout: 500 * time.Millisecond}

	for i := 0; i < 10; i++ {
		log.Printf("checking if service up, try=%d\n", i)

		res, err := c.Get("http://localhost:8080/health/readiness")
		if err == nil && res.StatusCode == 200 {
			return
		}

		time.Sleep(1 * time.Second)
	}

	panic("could not verify service has started")
}

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	mux.HandleFunc("/status.xml", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	ctx, cancelCtx := context.WithCancel(context.Background())

	app := cli.App{
		Commands: []*cli.Command{
			server.Cmd,
		},
	}

	dsn := generateSqliteInMemoryDsn()

	u, _ := url.Parse(srv.URL)
	os.Setenv("TVHGO_TVHEADEND_HOST", "127.0.0.1")
	os.Setenv("TVHGO_TVHEADEND_PORT", u.Port())
	os.Setenv("TVHGO_DATABASE_PATH", dsn)

	go app.RunContext(ctx, []string{"", "server"})

	waitForServiceIsReady()

	addTestData(dsn)

	code := m.Run()

	time.Sleep(10 * time.Second)
	cancelCtx()

	os.Exit(code)
}
