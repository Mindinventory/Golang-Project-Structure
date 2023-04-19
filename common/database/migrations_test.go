package database

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrations_Sort(t *testing.T) {
	migrations := Migrations{
		{Path: "/foo/bar/baz/V20__create_things.sql"},
		{Path: "/foo/bar/baz/V5__create_things.sql"},
		{Path: "/foo/bar/baz/V123__create_things.sql"},
	}

	sort.Sort(migrations)

	assert.Equal(t, 5, migrations[0].Version())
	assert.Equal(t, 20, migrations[1].Version())
	assert.Equal(t, 123, migrations[2].Version())
}

func TestMigration_Filename(t *testing.T) {
	m := Migration{Path: "/foo/bar/baz/V1__create_things.sql"}
	assert.Equal(t, "V1__create_things.sql", m.Filename())
}

func TestMigration_Name(t *testing.T) {
	m := Migration{Path: "/foo/bar/baz/V1__create_things.sql"}
	assert.Equal(t, "create_things", m.Name())
}

func TestMigration_SQL(t *testing.T) {
	sql := "CREATE TABLE things;"
	m := Migration{Data: []byte(sql)}
	assert.Equal(t, sql, m.SQL())
}

func TestMigration_Version(t *testing.T) {
	tests := []struct {
		expected int
		path     string
	}{
		{-1, "/foo/bar/baz/no_number.sql"},
		{1, "/foo/bar/baz/V1__create_things.sql"},
		{10, "/foo/bar/baz/V10__create_things.sql"},
		{123, "/foo/bar/baz/V123__create_things.sql"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			migration := &Migration{Path: tt.path}
			assert.Equal(t, tt.expected, migration.Version())
		})
	}
}
