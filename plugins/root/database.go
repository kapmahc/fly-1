package root

import (
	"database/sql"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/steinbacher/goose"
)

// DatabaseMigrate migrate database
func DatabaseMigrate() error {
	conf, err := dbConf()
	if err != nil {
		return err
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		return err
	}

	return goose.RunMigrations(conf, conf.MigrationsDir, target)
}

// DatabaseVersion data version
func DatabaseVersion(wrt io.Writer) error {
	conf, err := dbConf()
	if err != nil {
		return err
	}

	// collect all migrations
	migrations, err := goose.CollectMigrations(conf.MigrationsDir)
	if err != nil {
		return err
	}

	db, err := goose.OpenDBFromDBConf(conf)
	if err != nil {
		return err
	}
	defer db.Close()

	// must ensure that the version table exists if we're running on a pristine DB
	if _, err = goose.EnsureDBVersion(conf, db); err != nil {
		return err
	}

	fmt.Fprintln(wrt, "    Applied At                  Migration")
	fmt.Fprintln(wrt, "    =======================================")
	for _, m := range migrations {
		if err = printMigrationsStatus(wrt, db, m.Version, filepath.Base(m.Source)); err != nil {
			return err
		}
	}
	return nil
}

func printMigrationsStatus(wrt io.Writer, db *sql.DB, version int64, script string) error {
	var row goose.Migration
	q := fmt.Sprintf("SELECT tstamp, is_applied FROM goose_db_version WHERE version_id=%d ORDER BY tstamp DESC LIMIT 1", version)
	e := db.QueryRow(q).Scan(&row.TStamp, &row.IsApplied)

	if e != nil && e != sql.ErrNoRows {
		return e
	}

	var appliedAt string

	if row.IsApplied {
		appliedAt = row.TStamp.Format(time.ANSIC)
	} else {
		appliedAt = "Pending"
	}

	fmt.Fprintf(wrt, "    %-24s -- %v\n", appliedAt, script)
	return nil
}

func dbConf() (*goose.DBConf, error) {
	drv := goose.DBDriver{
		Name: beego.AppConfig.String("databasedriver"),
		DSN:  beego.AppConfig.String("databasesource"),
	}
	switch drv.Name {
	case "postgres":
		drv.Import = "github.com/lib/pq"
		drv.Dialect = &goose.PostgresDialect{}
	case "mysql":
		drv.Import = "github.com/go-sql-driver/mysql"
		drv.Dialect = &goose.MySqlDialect{}
	default:
		return nil, fmt.Errorf("unsupported driver %s", drv.Name)
	}
	return &goose.DBConf{
		Driver:        drv,
		MigrationsDir: path.Join("db", drv.Name, "migrations"),
	}, nil
}
