package main

import (
	"gopkg.in/gcfg.v1"
	"io/ioutil"
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
# A regular MySQL DSN (eg: user:password@tcp(127.0.0.1:3306)/mysql)
# Ideally this user would only have 'REPLICATION CLIENT'
dsn = 

[notify]
# You can have as many emails variables as you want, they'll all be used
emails = 

[smtp]
# A SMTP server (eg: example.com)
server = 
# A SMTP port (eg: 25)
port = 
# A from email address (eg: example@example.com)
from = 
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

	Smtp struct {
		Server string
		Port   int
		From   string
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

// Creates the config.gcfg if it does not exist.
func (c *Configuration) ensureConfigExists(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return ioutil.WriteFile(file, []byte(CONFIG_EXAMPLE), 0644)
	} else {
		return nil
	}
}
