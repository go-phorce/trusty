package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/ekspand/trusty/internal/db/model"
	"github.com/ekspand/trusty/internal/db/pgsql"
	"github.com/go-phorce/dolly/fileutil"
	"github.com/go-phorce/dolly/xlog"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/juju/errors"

	// register Postgres driver
	_ "github.com/lib/pq"

	// register file driver for migration
	_ "github.com/golang-migrate/migrate/source/file"
)

var logger = xlog.NewPackageLogger("github.com/ekspand/trusty/internal", "db")

// IDGenerator defines an interface to generate unique ID accross the cluster
type IDGenerator interface {
	// NextID generates a next unique ID.
	NextID() (uint64, error)
}

// UsersDb defines an interface for CRUD operations on Users and Teams
type UsersDb interface {
	// LoginUser returns User
	LoginUser(ctx context.Context, user *model.User) (*model.User, error)
	// GetUser returns User
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

// OrgsDb defines an interface for CRUD operations on Orgs
type OrgsDb interface {
	// UpdateOrg inserts or updates Organization
	UpdateOrg(ctx context.Context, org *model.Organization) (*model.Organization, error)
	// GetOrg returns Organization
	GetOrg(ctx context.Context, id int64) (*model.Organization, error)
	// RemoveOrg deletes org and all its members
	RemoveOrg(ctx context.Context, id int64) error

	// UpdateRepo inserts or updates Repository
	UpdateRepo(ctx context.Context, repo *model.Repository) (*model.Repository, error)
	// GetRepo returns Repository
	GetRepo(ctx context.Context, id int64) (*model.Repository, error)
	// TODO: RemoveRepo

	// AddOrgMember adds a user to Org
	AddOrgMember(ctx context.Context, orgID, userID int64, role, membershipSource string) (*model.OrgMembership, error)
	// GetOrgMembers returns list of membership info
	GetOrgMembers(ctx context.Context, orgID int64) ([]*model.OrgMemberInfo, error)
	// RemoveOrgMembers removes users from the org
	RemoveOrgMembers(ctx context.Context, orgID int64, all bool) ([]*model.OrgMembership, error)
	// RemoveOrgMember remove users from the org
	RemoveOrgMember(ctx context.Context, orgID, memberID int64) (*model.OrgMembership, error)
	// GetUserMemberships returns list of membership info
	GetUserMemberships(ctx context.Context, userID int64) ([]*model.OrgMemberInfo, error)
	// GetUserOrgs returns list of orgs
	GetUserOrgs(ctx context.Context, userID int64) ([]*model.Organization, error)
}

// RaDb defines an interface for Registration Authority
type RaDb interface {
	// RegisterRootCertificate registers Root Cert
	RegisterRootCertificate(ctx context.Context, crt *model.RootCertificate) (*model.RootCertificate, error)
	// RemoveRootCertificate removes Root Cert
	RemoveRootCertificate(ctx context.Context, id int64) error
	// GetRootCertificates returns list of Root certs
	GetRootCertificates(ctx context.Context, ownerID int64) ([]*model.RootCertificate, error)
}

// Provider represents SQL client instance
type Provider interface {
	UsersDb
	OrgsDb
	RaDb

	// DB returns underlying DB connection
	DB() *sql.DB
	// Close connection and release resources
	Close() (err error)
	// NextID returns unique ID
	NextID() (uint64, error)
}

// Migrate performs the db migration
func Migrate(migrationsDir string, db *sql.DB) error {
	logger.Tracef("src=Migrate, reason=load, directory=%q", migrationsDir)
	if _, err := os.Stat(migrationsDir); err != nil {
		return errors.Annotatef(err, "directory %q inaccessible", migrationsDir)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Trace(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsDir),
		"postgres",
		driver)
	if err != nil {
		return errors.Trace(err)
	}

	version, _, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return errors.Trace(err)
	}
	if err == migrate.ErrNilVersion {
		logger.Tracef("src=Migrate, reason=initial_state, version=nil")
	} else {
		logger.Tracef("src=Migrate, reason=initial_state, version=%d", version)
	}

	err = m.Up()
	if err != nil {
		return errors.Trace(err)
	}

	version, _, err = m.Version()
	if err != nil {
		return errors.Trace(err)
	}
	logger.Infof("src=Migrate, reason=current_state, version=%d", version)

	return nil
}

// New creates a Provider instance
func New(driverName, dataSourceName, migrationsDir string, nextID func() (uint64, error)) (Provider, error) {
	ds, err := fileutil.LoadConfigWithSchema(dataSourceName)
	if err != nil {
		return nil, errors.Trace(err)
	}

	db, err := sql.Open(driverName, ds)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = Migrate(migrationsDir, db)
	if err != nil && !strings.Contains(err.Error(), "no change") {
		return nil, errors.Trace(err)
	}

	return pgsql.New(db, nextID)
}
