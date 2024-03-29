package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	views "github.com/neuralknight/backend-views"
	log "github.com/sirupsen/logrus"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type NKnightSuite struct {
	srv      *httptest.Server
	client   *http.Client
	endpoint string
}

var _ = Suite(&NKnightSuite{})

func (s *NKnightSuite) TestCmdLoop(c *C) {
	defer func() {
		switch err := recover().(type) {
		case error:
			log.Println(err.Error())
		default:
			break
		}
	}()
	CmdLoop(s.endpoint)
}

func (s *NKnightSuite) TestMainEntry(c *C) {
	defer func() {
		switch err := recover().(type) {
		case error:
			log.Println(err.Error())
		default:
			break
		}
	}()
	Main()
}

func (s *NKnightSuite) SetUpSuite(c *C) {
	s.srv = httptest.NewServer(views.Handler{})
	s.client = s.srv.Client()
	s.endpoint = s.srv.URL
}

func (s *NKnightSuite) SetUpTest(c *C) {}

func (s *NKnightSuite) TearDownTest(c *C) {
	db, _ := gorm.Open("sqlite3", "chess.db")
	db = db.Begin()
	defer db.Commit()
	db.DropTableIfExists("board_models", "agent_models")
}

func (s *NKnightSuite) TearDownSuite(c *C) {
	s.srv.Close()
}
