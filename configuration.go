package main

import (
	"io/ioutil"
	"gopkg.in/gcfg.v1"
	"os"
)

var (
	DefaultConfigFile = "config.gcfg"
)

const (
	// The example file kept in version control. We'll copy and load from this
	// by default.
	CONFIG_EXAMPLE = `; mysql-replication-monior
[mysql-server "slave1"]
# A regular MySQL DSN (eg user:password@tcp(127.0.0.1:3306)/mysql)
# Ideally this user would only have 'REPLICATION CLIENT'
dsn =

[notify]
emails[] = luke@axxim.net
`
)

type MysqlServer struct {
	Dsn string
}

type Configuration struct {
	Mysql_Server map[string]*struct {
		Dsn string
	}

	Notify struct {
		Emails []string
	}
}

// Reads the configuration from the config file, copying a config into
// place from the example if one does not yet exist.
func (c *Configuration) load(file string) error {
	err := c.ensureConfigExists(file)
	if err != nil {
		return err
	}

	return gcfg.ReadFileInto(c, file)
}

// Creates the config.json if it does not exist.
func (c *Configuration) ensureConfigExists(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return ioutil.WriteFile(file, []byte(CONFIG_EXAMPLE), 0644)
	} else {
		return nil
	}
}