package gormdb

import (
	"github.com/kuroshibaz/config"
	"testing"
)

var conn, connErr = ConnectSQL(&config.Database{
	Host:     "localhost",
	Name:     "kuroshibaz",
	Username: "postgres",
	Password: "postgres",
	Port:     5432,
	SSLMode:  false,
	Debug:    true,
	Timezone: "Asia/Bangkok",
})

func TestDB_AssignRoleToUser(t *testing.T) {
	if connErr != nil {
		t.Error(connErr)
	}

	result, err := conn.AssignRoleToUser(1, 1)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestDB_AssignPermissionToRole(t *testing.T) {
	conn.AssignPermissionToRole("Admin")
}

func TestDB_CountViews(t *testing.T) {
	v, err := conn.CountViews("blog-2")
	if err != nil {
		t.Error(err)
	}
	t.Log("view count: ", v)
}
