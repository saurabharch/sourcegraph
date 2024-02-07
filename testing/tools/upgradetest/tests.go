package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"

	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/run"

	"github.com/sourcegraph/sourcegraph/lib/errors"
)

// Tests are the basic unit of this program and represent a version being tested. A tests methods are generally used to control its logging behavior.
// Tests are further organized by TestResults, a test result aggregation.
type Test struct {
	Version  semver.Version
	Type     string
	Runtime  time.Duration
	LogLines []string
	Errors   []error
}

// Addlog registers a log entry.
func (t *Test) AddLog(log string) {
	t.LogLines = append(t.LogLines, log)
}

// AddError registers an error.
func (t *Test) AddError(err error) {
	t.LogLines = append(t.LogLines, err.Error())
	t.Errors = append(t.Errors, err)
}

// DisplayErrors prints errors to stdout
func (t *Test) DisplayErrors() {
	for _, err := range t.Errors {
		fmt.Println(err)
	}
}

// DisplayLog prints logs to stdout
func (t *Test) DisplayLog() {
	for _, log := range t.LogLines {
		fmt.Println(log)
	}
}

// TestResults is a collection of tests, organized by type. Its methods are generally used to control its logging behavior.
type TestResults struct {
	StandardUpgradeTests []Test
	MVUUpgradeTests      []Test
	AutoupgradeTests     []Test
	Mutex                sync.Mutex
}

// AddStdTest adds a standard test to the results
func (r *TestResults) AddStdTest(test Test) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.StandardUpgradeTests = append(r.StandardUpgradeTests, test)
}

// AddMVUTest adds a multiversion upgrade test to the results
func (r *TestResults) AddMVUTest(test Test) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.MVUUpgradeTests = append(r.MVUUpgradeTests, test)
}

// AddAutoTest adds an autoupgrade test to the results
func (r *TestResults) AddAutoTest(test Test) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.AutoupgradeTests = append(r.AutoupgradeTests, test)
}

// Used in all-type test
type typeVersion struct {
	Type    string
	Version *semver.Version
}

// Known bug versions
// versions 4.1.0 to v4.4.2 are affected by a known bug in MVU if initialized in these versions: https://github.com/sourcegraph/sourcegraph/pull/46969
var knownBugVersions = []string{
	"4.1.0",
	"4.1.1",
	"4.1.2",
	"4.1.3",
	"4.2.0",
	"4.2.1",
	"4.3.0",
	"4.3.1",
	"4.4.0",
	"4.4.1",
	"4.4.2",
}

// PrintSimpleResults prints a quick view of test results, on an errored test only the first line of the error is printed.
func (r *TestResults) PrintSimpleResults() {
	if len(r.StandardUpgradeTests) != 0 {
		stdRes := []string{}
		for _, test := range r.StandardUpgradeTests {
			if 0 < len(test.Errors) {
				stdRes = append(stdRes, fmt.Sprintf("🚨 %s Failed -- %s\n%s", test.Version.String(), test.Runtime, test.Errors[0]))
			} else {
				stdRes = append(stdRes, fmt.Sprintf("✅ %s Passed -- %s ", test.Version.String(), test.Runtime))
			}
		}
		fmt.Println("--- 🕵️  Standard Upgrade Tests:")
		fmt.Println(strings.Join(stdRes, "\n"))
	}
	if len(r.MVUUpgradeTests) != 0 {
		mvuRes := []string{}
		for _, test := range r.MVUUpgradeTests {
			if 0 < len(test.Errors) {
				mvuRes = append(mvuRes, fmt.Sprintf("🚨 %s Failed -- %s\n%s", test.Version.String(), test.Runtime, test.Errors[0]))
			} else {
				mvuRes = append(mvuRes, fmt.Sprintf("✅ %s Passed -- %s", test.Version.String(), test.Runtime))
			}
		}
		fmt.Println("--- 🕵️  Multiversion Upgrade Tests:")
		fmt.Println(strings.Join(mvuRes, "\n"))
	}
	if len(r.AutoupgradeTests) != 0 {
		autoRes := []string{}
		for _, test := range r.AutoupgradeTests {
			if 0 < len(test.Errors) {
				autoRes = append(autoRes, fmt.Sprintf("🚨 %s Failed -- %s\n%s", test.Version.String(), test.Runtime, test.Errors[0]))
			} else {
				autoRes = append(autoRes, fmt.Sprintf("✅ %s Passed -- %s", test.Version.String(), test.Runtime))
			}
		}
		fmt.Println("--- 🕵️  Auto Upgrade Tests:")
		fmt.Println(strings.Join(autoRes, "\n"))
	}
}

// DisplayErrrors, prints errors for all tests that errored.
func (r *TestResults) DisplayErrors() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	for _, test := range r.StandardUpgradeTests {
		if 0 < len(test.Errors) {
			fmt.Printf("--- 🚨 Standard Upgrade Test %s Failed:\n", test.Version.String())
			test.DisplayErrors()
		}
	}
	for _, test := range r.MVUUpgradeTests {
		if 0 < len(test.Errors) {
			fmt.Printf("--- 🚨 Multiversion Upgrade Test %s Failed:\n", test.Version.String())
			test.DisplayErrors()
		}
	}
	for _, test := range r.AutoupgradeTests {
		if 0 < len(test.Errors) {
			fmt.Printf("--- 🚨 Auto Upgrade Test %s Failed:\n", test.Version.String())
			test.DisplayErrors()
		}
	}
}

// OrderByVersion orders tests TestResults by their test.Version value
func (r *TestResults) OrderByVersion() {
	sort.Slice(r.StandardUpgradeTests, func(i, j int) bool {
		return r.StandardUpgradeTests[i].Version.LessThan(&r.StandardUpgradeTests[j].Version)
	})
	sort.Slice(r.MVUUpgradeTests, func(i, j int) bool {
		return r.MVUUpgradeTests[i].Version.LessThan(&r.MVUUpgradeTests[j].Version)
	})
	sort.Slice(r.AutoupgradeTests, func(i, j int) bool {
		return r.AutoupgradeTests[i].Version.LessThan(&r.AutoupgradeTests[j].Version)
	})
}

// standardUpgradeTest initializes Sourcegraph's dbs and runs a standard upgrade
// i.e. an upgrade test between some last minor version and the current release candidate
func standardUpgradeTest(ctx context.Context, initVersion, targetVersion, latestStableVersion *semver.Version) Test {
	//start test env
	test, networkName, dbs, cleanup, err := setupTestEnv(ctx, "standard", initVersion)
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to setup env: %w", err))
		cleanup()
		return test
	}
	defer cleanup()

	// ensure env correctly initialized
	if err := validateDBs(ctx, &test, initVersion.String(), fmt.Sprintf("sourcegraph/migrator:%s", latestStableVersion.String()), networkName, dbs, false); err != nil {
		test.AddError(errors.Newf("🚨 Upgrade failed: %w", err))
		return test
	}

	test.AddLog("-- ⚙️  performing standard upgrade")

	// Run standard upgrade via migrators "up" command
	out, err := run.Cmd(ctx, dockerMigratorBaseString(test, "up", "migrator:candidate", networkName, dbs)...).Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to upgrade: %w", err))
		cleanup()
		return test
	}
	test.AddLog(out)

	fctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start frontend with candidate
	var cleanFrontend func()
	cleanFrontend, err = startFrontend(fctx, test, "frontend", "candidate", networkName, false, dbs)
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to start candidate frontend: %w", err))
		cleanFrontend()
		return test
	}
	defer cleanFrontend()

	test.AddLog("-- ⚙️  post upgrade validation")
	// Validate the upgrade
	if err := validateDBs(ctx, &test, targetVersion.String(), "migrator:candidate", networkName, dbs, true); err != nil {
		test.AddError(errors.Newf("🚨 Upgrade failed: %w", err))
		return test
	}

	return test
}

// multiversionUpgradeTest tests the migrator upgrade command,
// initializing the three main dbs and conducting an upgrade to the release candidate version
func multiversionUpgradeTest(ctx context.Context, initVersion, targetVersion, latestStableVersion *semver.Version) Test {
	test, networkName, dbs, cleanup, err := setupTestEnv(ctx, "multiversion", initVersion)
	if err != nil {
		fmt.Println("🚨 failed to setup env: ", err)
		cleanup()
		return test
	}
	defer cleanup()

	// ensure env correctly initialized, always use latest migrator for drift check,
	// this allows us to avoid issues from changes in migrators invocation
	if err := validateDBs(ctx, &test, initVersion.String(), fmt.Sprintf("sourcegraph/migrator:%s", latestStableVersion.String()), networkName, dbs, false); err != nil {
		test.AddError(errors.Newf("🚨 Initializing env in multiversion test failed: %w", err))
		return test
	}

	// Run multiversion upgrade using candidate image
	//
	// If the build version in the test has been stamped this is the "to" input for the upgrade command. If the builds arent stamped we use the latest stable release verstion,
	// as the target for `migrator upgrade`. This is safe because we assume the target version of the test is always the latest release version. Therefore `migrator up` should always be a single minor version away.
	var toVersion string
	if targetVersion.String() != "0.0.0+dev" { // if version is stamped
		toVersion = targetVersion.String()
	} else {
		toVersion = latestStableVersion.String()
	}
	test.AddLog(fmt.Sprintf("-- ⚙️  performing multiversion upgrade (--from %s --to %s)", initVersion.String(), toVersion))
	out, err := run.Cmd(ctx,
		dockerMigratorBaseString(test, fmt.Sprintf("upgrade --from %s --to %s --ignore-migrator-update", initVersion.String(), toVersion), "migrator:candidate", networkName, dbs)...).
		Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to upgrade: %w", err))
		cleanup()
		return test
	}
	test.AddLog(out)

	// Run migrator up with migrator candidate to apply any patch migrations defined on the candidate version
	out, err = run.Cmd(ctx,
		dockerMigratorBaseString(test, "up", "migrator:candidate", networkName, dbs)...).
		Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to upgrade: %w", err))
		cleanup()
		return test
	}
	test.AddLog(out)

	// Start frontend with candidate
	var cleanFrontend func()
	cleanFrontend, err = startFrontend(ctx, test, "frontend", "candidate", networkName, false, dbs)
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to start candidate frontend: %w", err))
		cleanFrontend()
		return test
	}
	defer cleanFrontend()

	test.AddLog("-- ⚙️  post upgrade validation")
	// Validate the upgrade
	if err := validateDBs(ctx, &test, targetVersion.String(), "migrator:candidate", networkName, dbs, true); err != nil {
		test.AddError(errors.Newf("🚨 Upgrade failed: %w", err))
		return test
	}

	return test
}

// Logic in tryAutoUpgrade is not compatible with dev builds. The autoUpgrade test must be run with a stamp version
//
// Without this in place autoupgrade fails and exits while trying to make an oobmigration comparison here: https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/cmd/frontend/internal/cli/autoupgrade.go?L67-76
// {"SeverityText":"WARN","Timestamp":1706721478276103721,"InstrumentationScope":"frontend","Caller":"cli/autoupgrade.go:73","Function":"github.com/sourcegraph/sourcegraph/cmd/frontend/internal/cli.tryAutoUpgrade","Body":"unexpected string for desired instance schema version, skipping auto-upgrade","Resource":{"service.name":"frontend","service.version":"devVersion","service.instance.id":"487754e1c54a"},"Attributes":{"version":"devVersion"}}
func autoUpgradeTest(ctx context.Context, initVersion, targetVersion, latestStableVersion *semver.Version) Test {
	//start test env
	test, networkName, dbs, cleanup, err := setupTestEnv(ctx, "auto", initVersion)
	if err != nil {
		test.AddError(errors.Newf("failed to setup env: %w", err))
		cleanup()
		return test
	}
	defer cleanup()

	// ensure env correctly initialized, always use latest migrator for drift check,
	// this allows us to avoid issues from changes in migrators invocation
	if err := validateDBs(ctx, &test, initVersion.String(), fmt.Sprintf("sourcegraph/migrator:%s", latestStableVersion.String()), networkName, dbs, false); err != nil {
		test.AddError(errors.Newf("🚨 Initializing env in multiversion test failed: %w", err))
		return test
	}

	// Set SRC_AUTOUPGRADE=true on Migrator and Frontend containers. Then start the frontend container.
	test.AddLog("-- ⚙️  performing auto upgrade")

	fctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start frontend with candidate
	var cleanFrontend func()
	cleanFrontend, err = startFrontend(fctx, test, "frontend", "candidate", networkName, true, dbs)
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to start candidate frontend: %w", err))
		cleanFrontend()
		return test
	}
	defer cleanFrontend()

	test.AddLog("-- ⚙️  post upgrade validation")
	// Validate the upgrade
	if err := validateDBs(ctx, &test, targetVersion.String(), "migrator:candidate", networkName, dbs, true); err != nil {
		test.AddError(errors.Newf("🚨 Upgrade failed: %w", err))
		return test
	}

	return test
}

// testDB is an organizational type to make orchestrating the three dbs easier, and also to store a dynamically allocated port for postgres
type testDB struct {
	DbName            string
	ContainerName     string
	Image             string
	ContainerHostPort string
}

// setupTestEnv initializeses a test environment and object. Creates a docker network for testing as well as instances of our three databases. Returning a cleanup function.
// An instance of Sourcegraph-Frontend is also started to initialize the versions table of the database.
// TODO: setupTestEnv should seed some initial data at the target initVersion. This will be usefull for testing OOB migrations
func setupTestEnv(ctx context.Context, testType string, initVersion *semver.Version) (test Test, networkName string, dbs []*testDB, cleanup func(), err error) {
	test = Test{
		Version:  *initVersion,
		Type:     testType,
		LogLines: []string{},
		Errors:   []error{},
	}

	if testType == "standard" {
		test.AddLog("--- 🕵️  standard upgrade test")
	}
	if testType == "multiversion" {
		test.AddLog("--- 🕵️  multiversion upgrade test")
	}
	if testType == "auto" {
		test.AddLog("--- 🕵️  auto upgrade test")
	}
	test.AddLog(fmt.Sprintf("Upgrading from version (%s) to release candidate.", initVersion))
	test.AddLog("-- 🏗️  setting up test environment")

	// Create a docker network for testing
	//
	// Docker bridge networks take up a lot of the docker daemons available port allocation. We run only a limited amount of test parallelization to get around this.
	// see https://straz.to/2021-09-08-docker-address-pools/
	networkName = fmt.Sprintf("%s_test_%s", testType, initVersion)
	test.AddLog(fmt.Sprintf("🐋 creating network %s", networkName))

	out, err := run.Cmd(ctx, "docker", "network", "create", networkName).Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to create test network: %s", err))
	}
	test.AddLog(out)

	// Note that we changed postgres versions in very early versions of Sourcegraph,
	// In v3.38+ we use image postgres-12-alpine,
	// in v3.37-v3.30 we use postgres-12.6-alpine,
	// in v3.29-v3.27 we use postgres-12.6
	// in v3.26 and earliar we use postgres:11.4
	//
	// This isn't relevant since this test will only ever initialize instances v3.38+
	// worth noting in case this changes in the future.
	dbs = []*testDB{
		{"pgsql", fmt.Sprintf("%s_pgsql_%s", testType, initVersion), "postgres-12-alpine", ""},
		{"codeintel-db", fmt.Sprintf("%s_codeintel-db_%s", testType, initVersion), "codeintel-db", ""},
		{"codeinsights-db", fmt.Sprintf("%s_codeinsights-db_%s", testType, initVersion), "codeinsights-db", ""},
	}

	// Here we create the three databases using docker run.
	for _, db := range dbs {
		test.AddLog(fmt.Sprintf("🐋 creating %s, with db image %s:%s", db.ContainerName, db.Image, initVersion))
		err := run.Cmd(ctx, "docker", "run", "--rm",
			"--detach",
			"--platform", "linux/amd64",
			"--name", db.ContainerName,
			"--network", networkName,
			"-p", "5432",
			fmt.Sprintf("sourcegraph/%s:%s", db.Image, initVersion),
		).Run().Wait()
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to create test databases: %s", err))
		}
		// get the dynamically allocated port and register it to the test
		port, err := run.Cmd(ctx, "docker", "port", db.ContainerName, "5432").Run().String()
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to get port for %s: %s", db.ContainerName, err))
		}
		db.ContainerHostPort = port
	}

	// Create a timeout to validate the databases have initialized, this is to prevent a hung test
	// When many goroutines are running this test this is a point of failure.
	dbPingTimeout, cancel := context.WithTimeout(ctx, time.Second*120)
	wgDbPing := pool.New().WithErrors().WithContext(dbPingTimeout)
	defer cancel()

	// Here we poll/ping the dbs to ensure postgres has initialized before we make calls to the databases.
	for _, db := range dbs {
		db := db // this closure locks the index for the inner for loop
		wgDbPing.Go(func(ctx context.Context) error {
			dbClient, err := sql.Open("postgres", fmt.Sprintf("postgres://sg@%s/sg?sslmode=disable", db.ContainerHostPort))
			if err != nil {
				test.AddError(errors.Newf("🚨 failed to connect to %s: %s", db.DbName, err))
			}
			defer dbClient.Close()
			for {
				select {
				case <-dbPingTimeout.Done():
					return dbPingTimeout.Err()
				default:
				}
				err = dbClient.Ping()
				if err != nil {
					test.AddLog(fmt.Sprintf(" ... pinging %s", db.DbName))
					if err == sql.ErrConnDone || strings.Contains(err.Error(), "connection refused") {
						test.AddError(errors.Newf("🚨 unrecoverable error pinging %s: %w", db.DbName, err))
						return err
					}
					time.Sleep(1 * time.Second)
					continue
				} else {
					test.AddLog(fmt.Sprintf("✅ %s is up", db.DbName))
					return nil
				}
			}
		})
	}
	if err := wgDbPing.Wait(); err != nil {
		test.AddError(errors.Newf("🚨 containerized database startup error: %w", err))
	}

	// Initialize the databases by running migrator with the `up` command.
	test.LogLines = append(test.LogLines, "-- 🏗️  initializing database schemas with migrator")
	out, err = run.Cmd(ctx, dockerMigratorBaseString(test, "up", fmt.Sprintf("sourcegraph/migrator:%s", initVersion), networkName, dbs)...).Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to initialize database: %w", err))
	}
	test.AddLog(out)

	// Verify that the databases are initialized.
	test.AddLog("🔎 checking db schemas initialized")
	for _, db := range dbs {
		dbClient, err := sql.Open("postgres", fmt.Sprintf("postgres://sg@%s/sg?sslmode=disable", db.ContainerHostPort))
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to connect to %s: %s", db.DbName, err))
			continue
		}
		defer dbClient.Close()

		// check if tables have been created
		rows, err := dbClient.Query(`SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='public';`)
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to check %s for init: %s", db.DbName, err))
			continue
		}
		defer rows.Close()
		if err := rows.Err(); err != nil {
			test.AddError(errors.Newf("🚨 failed to check %s for init: %s", db.DbName, err))
			continue
		} else {
			test.AddLog(fmt.Sprintf("✅ %s initialized", db.DbName))
		}
	}

	//start frontend and poll db until initial version is set by frontend
	var cleanFrontend func()
	cleanFrontend, err = startFrontend(ctx, test, "sourcegraph/frontend", initVersion.String(), networkName, false, dbs)
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to start frontend: %w", err))
	}
	defer cleanFrontend()

	// Return a cleanup function that will remove the containers and network.
	cleanup = func() {
		test.LogLines = append(test.LogLines, "🧹 removing database containers")
		out, err := run.Cmd(ctx, "docker", "container", "stop",
			dbs[0].ContainerName,
			dbs[1].ContainerName,
			dbs[2].ContainerName).
			Run().String()
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to stop database containers after testing: %w", err))
		}
		test.AddLog(out)
		out, err = run.Cmd(ctx, "docker", "container", "rm",
			dbs[0].ContainerName,
			dbs[1].ContainerName,
			dbs[2].ContainerName).
			Run().String()
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to remove database containers after testing: %w", err))
		}
		test.AddLog(out)
		test.AddLog("🧹 removing testing network")
		out, err = run.Cmd(ctx, "docker", "network", "rm", networkName).Run().String()
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to remove test network after testing: %w", err))
		}
		test.AddLog(out)
	}

	test.AddLog("-- 🏗️  setup complete")

	return test, networkName, dbs, cleanup, err
}

// validateDBs runs a few tests to assess the readiness of the database and whether or not drift exists on the schema.
// It is used in initializing a new db as well as "validating" the db after an version change. This behavior is controlled by the upgrade parameter.
//
// TODO currently the post upgrade validation check relies on the latest commit on the branch for a schema version target, in version stamped tests we should target the stamp version schema.
func validateDBs(ctx context.Context, test *Test, version, migratorImage, networkName string, dbs []*testDB, postUpgrade bool) error {
	test.AddLog("-- ⚙️  validating dbs")

	// Get DB clients
	clients := make(map[string]*sql.DB)
	for _, db := range dbs {
		client, err := sql.Open("postgres", fmt.Sprintf("postgres://sg@%s/sg?sslmode=disable", db.ContainerHostPort))
		if err != nil {
			test.AddError(errors.Newf("🚨 failed to connect to %s: %w", db.DbName, err))
			return err
		}
		defer client.Close()

		clients[db.DbName] = client
	}

	// Verify the versions.version value in the frontend db
	test.AddLog("🔎 checking pgsql versions.version set")
	var testVersion string
	row := clients["pgsql"].QueryRowContext(ctx, `SELECT version FROM versions;`)
	if err := row.Scan(&testVersion); err != nil {
		test.AddError(errors.Newf("🚨 failed to get version from pgsql db: %w", err))
	}
	if version != testVersion {
		test.AddError(errors.Newf("🚨 versions.version not set: %s!= %s", version, testVersion))
	}

	test.AddLog(fmt.Sprintf("✅ versions.version set: %s", testVersion))

	// Check for any failed migrations in the migration logs tables
	// migration_logs table is introduced in v3.36.0
	for _, db := range dbs {
		test.AddLog(fmt.Sprintf("🔎 checking %s migration_logs", db.ContainerName))
		var numFailedMigrations int
		row = clients[db.DbName].QueryRowContext(ctx, `SELECT count(*) FROM migration_logs WHERE success=false;`)
		if err := row.Scan(&numFailedMigrations); err != nil {
			test.AddError(errors.Newf("🚨 failed to get failed migrations from %s db: %w", db.ContainerName, err))
		}
		if numFailedMigrations > 0 {
			test.AddError(errors.Newf("🚨 failed migrations found: %d", numFailedMigrations))
		}

		test.AddLog("✅ no failed migrations found")
	}

	// Check DBs for drift
	test.AddLog("🔎 Checking DBs for drift")
	if postUpgrade {
		// Get the last commit in the release branch, if validating an upgrade the upgrade boolean is true,
		// in this case the drift target is the latest commit on the release candidate branch.
		// If working on this, the drift check will fail if you have local commits not yet pushed to remote.
		// example schema check target: https://raw.githubusercontent.com/sourcegraph/sourcegraph/7648573357fef049e1a0bf11f400068ef83f2596/internal/database/schema.json
		var candidateGitHead bytes.Buffer
		if err := run.Cmd(ctx, "git", "rev-parse", "HEAD").Run().Stream(&candidateGitHead); err != nil {
			test.AddError(errors.Newf("🚨 failed to get latest commit on candidate branch: %w", err))
		}
		test.AddLog(fmt.Sprintf("Latest commit on candidate branch: %s", candidateGitHead.String()))
		for _, db := range dbs {
			out, err := run.Cmd(ctx, dockerMigratorBaseString(*test, fmt.Sprintf("drift --db %s --version %s --ignore-migrator-update --skip-version-check", db.DbName, candidateGitHead.String()),
				migratorImage, networkName, dbs)...).Run().String()
			if err != nil {
				test.AddError(errors.Newf("🚨 failed to check drift on %s: %s", db.DbName, err))
			}
			test.AddLog(out)
		}
	} else {
		for _, db := range dbs {
			out, err := run.Cmd(ctx, dockerMigratorBaseString(*test, fmt.Sprintf("drift --db %s --version v%s --ignore-migrator-update", db.DbName, version),
				migratorImage, networkName, dbs)...).Run().String()
			if err != nil {
				test.AddError(errors.Newf("🚨 failed to check drift on %s: %w", db.DbName, err))
			}
			test.AddLog(out)
		}
	}

	return nil
}

// startFrontend starts a frontend container and polls the pgsql database for certain states. When the state conditions are startFrontend returns a cleanup function that will stop and remove the frontend container.
// - checks for existence of site-config
// - checks that the version is set in pgsql
// - Optionally sets the auto upgrade env var to true or false.
func startFrontend(ctx context.Context, test Test, image, version, networkName string, auto bool, dbs []*testDB) (cleanup func(), err error) {
	stamp := ctx.Value(stampVersionKey{})

	hash, err := newContainerHash()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to get container hash: %w", err))
		return nil, err
	}
	test.AddLog(fmt.Sprintf("🐋 creating %s_frontend_%x", test.Type, hash))
	// define cleanup function to stop and remove the container
	cleanup = func() {
		test.AddLog("🧹 removing frontend container")
		out, err := run.Cmd(ctx, "docker", "container", "stop",
			fmt.Sprintf("%s_frontend_%x", test.Type, hash),
		).Run().String()
		if err != nil {
			fmt.Println("🚨 failed to stop frontend after testing: ", err)
		}
		test.AddLog(out)
		out, err = run.Cmd(ctx, "docker", "container", "rm",
			fmt.Sprintf("%s_frontend_%x", test.Type, hash),
		).Run().String()
		if err != nil {
			fmt.Println("🚨 failed to remove frontend after testing: ", err)
		}
		test.AddLog(out)
	}

	// construct docker command for running frontend container
	baseString := []string{
		"docker", "run",
		"--detach",
		"--platform", "linux/amd64",
		"--name", fmt.Sprintf("%s_frontend_%x", test.Type, hash),
	}
	envString := []string{
		"-e", "DEPLOY_TYPE=docker-container",
		"-e", fmt.Sprintf("PGHOST=%s", dbs[0].ContainerName),
		"-e", fmt.Sprintf("CODEINTEL_PGHOST=%s", dbs[1].ContainerName),
		"-e", fmt.Sprintf("CODEINSIGHTS_PGDATASOURCE=postgres://sg@%s:5432/sg?sslmode=disable", dbs[2].ContainerName),
	}
	if auto {
		envString = append(envString, "-e", "SRC_AUTOUPGRADE=true")
	}
	// ERROR
	// {"SeverityText":"FATAL","Timestamp":1706224238009644720,"InstrumentationScope":"sourcegraph","Caller":"svcmain/svcmain.go:167","Function":"github.com/sourcegraph/sourcegraph/internal/service/svcmain.run.func1","Body":"failed to start service","Resource":{"service.name":"frontend","service.version":"0.0.0+dev","service.instance.id":"79a3e3ca0bfc"},"Attributes":{"service":"frontend","error":"failed to connect to frontend database: database schema out of date"}}
	cmdString := []string{
		"--network", networkName,
		fmt.Sprintf("%s:%s", image, version),
	}
	baseString = append(baseString, envString...)
	cmdString = append(baseString, cmdString...)

	// Start the frontend container
	out, err := run.Cmd(ctx, cmdString...).Run().String()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to start frontend: %w", err))
		return cleanup, err
	}
	test.AddLog(fmt.Sprintf("frontend startup logs: %s", out))

	// poll db until initial versions.version is set
	setInitTimeout, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	test.AddLog("🔎 checking db initialization complete")

	dbClient, err := sql.Open("postgres", fmt.Sprintf("postgres://sg@%s/sg?sslmode=disable", dbs[0].ContainerHostPort))
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to connect to %s: %w", dbs[0].DbName, err))
	}
	defer dbClient.Close()

	// Poll till versions.version is set
	for {
		select {
		case <-ctx.Done():
			return cleanup, setInitTimeout.Err()
		case <-setInitTimeout.Done():
			return cleanup, setInitTimeout.Err()
		default:
		}
		// check version string set
		var dbVersion string
		row := dbClient.QueryRowContext(setInitTimeout, versionQuery)
		err = row.Scan(&dbVersion)
		if err != nil {
			test.AddLog(fmt.Sprintf("... querying versions.version: %s", err))
			time.Sleep(1 * time.Second)
			continue
		}
		// wait for the frontend to set the database versions.version value before considering the frontend startup complete.
		// "candidate" resolves to "0.0.0+dev" and should always be valid, so too a stamp version is valid
		if dbVersion == version || dbVersion == stamp || dbVersion == "0.0.0+dev" {
			test.AddLog(fmt.Sprintf("✅ versions.version is set: %s", dbVersion))
			break
		}
		if version != dbVersion {
			time.Sleep(1 * time.Second)
			test.AddLog(fmt.Sprintf(" ... waiting for versions.version to be set: %s", version))
			continue
		}
	}

	// poll db until site-config is initialized, migrator will sometimes fail if this initialization of the frontend db hasnt finished
	// returning an error like: "instance is new"
	for {
		select {
		case <-setInitTimeout.Done():
			return cleanup, setInitTimeout.Err()
		default:
		}
		// check version string set
		var siteConfig string
		row := dbClient.QueryRowContext(setInitTimeout, siteConfigQuery)
		err = row.Scan(&siteConfig)
		if err != nil {
			test.AddLog(fmt.Sprintf("... checking site-config initialized: %s", err))
			time.Sleep(1 * time.Second)
			continue
		}
		// Just test for site-config existence
		if siteConfig == "" {
			test.AddLog("... waiting for site-config to be initialized")
			time.Sleep(1 * time.Second)
			continue
		} else {
			test.AddLog("✅ site-config is initialized")
			break
		}
	}

	return cleanup, nil
}

const versionQuery = `SELECT version FROM versions;`

const siteConfigQuery = `
SELECT c.contents
FROM critical_and_site_config c
WHERE c.type = 'site'
ORDER BY c.id DESC
LIMIT 1
`

// dockerMigratorBaseString a slice of strings constituting the necessary arguments to run the migrator via docker container the CI test env.
func dockerMigratorBaseString(test Test, cmd, migratorImage, networkName string, dbs []*testDB) []string {
	hash, err := newContainerHash()
	if err != nil {
		test.AddError(errors.Newf("🚨 failed to get container hash: %w", err))
		return nil
	}
	baseString := []string{
		"docker", "run", "--rm",
		"--platform", "linux/amd64",
		"--name", fmt.Sprintf("%s_migrator_%x", test.Type, hash),
	}

	// This string construction will allow for possible env vars to be added later
	envString := []string{
		"-e", fmt.Sprintf("PGHOST=%s", dbs[0].ContainerName),
		"-e", "PGPORT=5432",
		"-e", "PGUSER=sg",
		"-e", "PGPASSWORD=sg",
		"-e", "PGDATABASE=sg",
		"-e", "PGSSLMODE=disable",
		"-e", fmt.Sprintf("CODEINTEL_PGHOST=%s", dbs[1].ContainerName),
		"-e", "CODEINTEL_PGPORT=5432",
		"-e", "CODEINTEL_PGUSER=sg",
		"-e", "CODEINTEL_PGPASSWORD=sg",
		"-e", "CODEINTEL_PGDATABASE=sg",
		"-e", "CODEINTEL_PGSSLMODE=disable",
		"-e", fmt.Sprintf("CODEINSIGHTS_PGHOST=%s", dbs[2].ContainerName),
		"-e", "CODEINSIGHTS_PGPORT=5432",
		"-e", "CODEINSIGHTS_PGUSER=sg", // starting codeinsights without frontend initializes with user sg rather than postgres
		"-e", "CODEINSIGHTS_PGPASSWORD=password",
		"-e", "CODEINSIGHTS_PGDATABASE=sg", // starting codeinsights without frontend initializes with database name as sg rather than postgres
		"-e", "CODEINSIGHTS_PGSSLMODE=disable",
		"-e", "SRC_LOG_LEVEL=debug",
	}

	cmdString := []string{
		"--network", networkName,
		migratorImage,
		cmd,
	}
	// append base string, env string, and cmd string and return the result
	baseString = append(baseString, envString...)
	return append(baseString, cmdString...)
}

// newContainerHash generates a random hash for naming containers in test, used for frontend and migrator
func newContainerHash() ([]byte, error) {
	hash := make([]byte, 4)
	_, err := rand.Read(hash)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// getVersions returns the latest minor semver version, as well as the latest full semver version.
//
// Technically MVU is supported v3.20 and forward, but in older versions codeinsights-db didnt exist and postgres was using version 11.4
// so we reduce the scope of the test, to cover only v3.39 and forward, for MVU and Auto upgrade testing.
func getVersions(cCtx *cli.Context) (latestMinor, latestFull, targetVersion *semver.Version, stdVersions, mvuVersions, autoVersions []*semver.Version, err error) {
	ctx := cCtx.Context

	tags, err := run.Cmd(ctx, "git",
		"for-each-ref",
		"--format", "'%(refname:short)'",
		"refs/tags").Run().Lines()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	var validTags []*semver.Version
	var latestMinorVer *semver.Version
	var latestFullVer *semver.Version

	// Get valid tags
	for _, tag := range tags {
		v, err := semver.NewVersion(tag)
		if err != nil {
			continue // skip non-matching tags
		}
		if v.Prerelease() != "" {
			continue // skip prereleases
		}
		// To simplify this testing range we'll only select a version tags from versions greater than v3.38
		// In v3.39 many things were normalized in the dbs:
		// - codeinsights-db moved from timescaleDB to posgres-12
		// - our image for codeintel-db and pgsql was normalized to postgres-12-alpine
		// - the migration_logs table exists, this was renamed from schema_migrations in v3.36.0
		// - migrator is introduced in v3.38.0
		if v.LessThan(semver.MustParse("v3.39.0")) {
			continue
		}
		validTags = append(validTags, v)
	}

	// Get latest Version and latestMinorVersion
	for _, tag := range validTags {
		// Track latest full version
		if latestFullVer == nil || tag.GreaterThan(latestFullVer) {
			latestFullVer = tag
		}
		latestMinorVer, err = semver.NewVersion(fmt.Sprintf("%d.%d.0", latestFullVer.Major(), latestFullVer.Minor()))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
	}

	// Get range for standardUpgrade test pool and MVU test pool
	for _, tag := range validTags {
		std, err := semver.NewConstraint(fmt.Sprintf(">= %d.%d.x", latestMinorVer.Major(), latestMinorVer.Minor()-1))
		if err != nil {
			fmt.Println("🚨 failed to collect versions for standard upgrade test: ", err)
		}
		// Sort versions into those for the standard test (one minor version behind the latest release candidate), and those for multiversion testing.
		if std.Check(tag) {
			stdVersions = append(stdVersions, tag)
		} else {
			mvuVersions = append(mvuVersions, tag)
		}
		// Auto upgrade tests cover all versions
		autoVersions = append(autoVersions, tag)
	}

	if latestMinorVer == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("No valid minor semver tags found")
	}
	if latestFullVer == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("No valid full semver tags found")
	}

	// Set target version to VERSION stamp if frontend and migrator are set at a stamped version, otherwise set it to 0.0.0+dev
	if cCtx.String("stamp-version") != "" {
		targetVersion = semver.MustParse(cCtx.String("stamp-version"))
	} else {
		targetVersion = semver.MustParse("0.0.0+dev") // If no stamp version is set, we assume version is in dev
	}

	return latestMinorVer, latestFullVer, targetVersion, stdVersions, mvuVersions, autoVersions, nil

}