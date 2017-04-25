package site

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/steinbacher/goose"
)

// GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	if err := p.dbMigrate(); err != nil {
		p.Abort(http.StatusInternalServerError)
	}
}

// PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {

}

// --------------------------------------------------------

func (p *Controller) dbMigrate() error {
	conf, err := p.dbConf()
	if err != nil {
		return err
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		return err
	}

	return goose.RunMigrations(conf, conf.MigrationsDir, target)
}

func (p *Controller) migrationsStatus(wrt io.Writer, db *sql.DB, version int64, script string) error {
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

func (p *Controller) dbConf() (*goose.DBConf, error) {
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
