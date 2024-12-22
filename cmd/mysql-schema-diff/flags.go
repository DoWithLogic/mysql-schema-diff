package main

import (
	"fmt"
	"strings"

	"github.com/DoWithLogic/mysql-schema-diff/pkg/log"
	"github.com/go-logfmt/logfmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

type connFlags struct {
	dsn string
}

type Config struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func createConnFlags(cmd *cobra.Command) *connFlags {
	flags := new(connFlags)

	// Don't mark dsn as a required flag.
	// Allow users to user the MYSQLHOST etc environment variables like `mysql`
	cmd.Flags().StringVar(&flags.dsn, "dsn", "", "Connection string for the database (DB password can be specified through MYSQLPASSWORD environment variable)")

	return flags
}

func parseConnConfig(c connFlags, logger log.Logger) (*mysql.Config, error) {
	if c.dsn == "" {
		logger.Warnf("DSN flag not set. Using libpq environment variables and default values.")
	}

	cfg, err := mysql.ParseDSN(c.dsn)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func logFmtToMap(logFmt string) (map[string]string, error) {
	logMap := make(map[string]string)
	decoder := logfmt.NewDecoder(strings.NewReader(logFmt))
	for decoder.ScanRecord() {
		if _, ok := logMap[string(decoder.Key())]; ok {
			return nil, fmt.Errorf("duplicate key %q in logfmt", string(decoder.Key()))
		}
		logMap[string(decoder.Key())] = string(decoder.Value())
	}

	if decoder.Err() != nil {
		return nil, decoder.Err()
	}

	return logMap, nil
}
