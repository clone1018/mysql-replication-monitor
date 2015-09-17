// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code was mostly taken from
// https://github.com/youtube/vitess/blob/master/go/vt/mysqlctl/replication.go
package main

import (
	"database/sql"
	"errors"
	"strconv"
)

var showSlaveStatusColumnNames = []string{
	"Slave_IO_State",
	"Master_Host",
	"Master_User",
	"Master_Port",
	"Connect_Retry",
	"Master_Log_File",
	"Read_Master_Log_Pos",
	"Relay_Log_File",
	"Relay_Log_Pos",
	"Relay_Master_Log_File",
	"Slave_IO_Running",
	"Slave_SQL_Running",
	"Replicate_Do_DB",
	"Replicate_Ignore_DB",
	"Replicate_Do_Table",
	"Replicate_Ignore_Table",
	"Replicate_Wild_Do_Table",
	"Replicate_Wild_Ignore_Table",
	"Last_Errno",
	"Last_Error",
	"Skip_Counter",
	"Exec_Master_Log_Pos",
	"Relay_Log_Space",
	"Until_Condition",
	"Until_Log_File",
	"Until_Log_Pos",
	"Master_SSL_Allowed",
	"Master_SSL_CA_File",
	"Master_SSL_CA_Path",
	"Master_SSL_Cert",
	"Master_SSL_Cipher",
	"Master_SSL_Key",
	"Seconds_Behind_Master",
	"Master_SSL_Verify_Server_Cert",
	"Last_IO_Errno",
	"Last_IO_Error",
	"Last_SQL_Errno",
	"Last_SQL_Error",
	"Exec_Master_Group_ID",
	"Connect_Using_Group_ID",
}

func SlaveStatus(db *sql.DB) (map[string]interface{}, error) {
	rows, err := db.Query("SHOW SLAVE STATUS")
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.New("Not a slave")
	}
	defer rows.Close()

	values := make([]interface{}, len(showSlaveStatusColumnNames))
	for i, _ := range values {
		var v sql.RawBytes
		values[i] = &v
	}

	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}

	slaveInfo := make(map[string]interface{})
	for i, name := range showSlaveStatusColumnNames {
		bp := values[i].(*sql.RawBytes)
		vs := string(*bp)
		vi, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			slaveInfo[name] = vs
		} else {
			slaveInfo[name] = vi
		}
	}

	return slaveInfo, nil
}
