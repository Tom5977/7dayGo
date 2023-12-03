package session_test

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/session"
	"os"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *session.Session {
	return session.New(TestDB, TestDial)
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
