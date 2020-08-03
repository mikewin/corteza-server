package test_store

import (
	"context"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/corteza/store/mysql"
	"github.com/cortezaproject/corteza-server/corteza/store/pgsql"
	"github.com/cortezaproject/corteza-server/tests/helpers"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

func init() {
	helpers.RecursiveDotEnvLoad()
}

func Test_Store(t *testing.T) {
	type (
		suite struct {
			name      string
			dsnEnvKey string
			init      func(ctx context.Context, dsn string) (store.Interface, error)
		}
	)

	var (
		ctx = context.Background()

		ss = []suite{
			{
				name:      "MySQL",
				dsnEnvKey: "RDBMS_MYSQL_DSN",
				init:      func(ctx context.Context, dsn string) (store.Interface, error) { return mysql.New(ctx, dsn) },
			},
			{
				name:      "PostgreSQL",
				dsnEnvKey: "RDBMS_PGSQL_DSN",
				init:      func(ctx context.Context, dsn string) (store.Interface, error) { return pgsql.New(ctx, dsn) },
			},
			{
				name:      "CockroachDB",
				dsnEnvKey: "RDBMS_COCKROACHDB_DSN",
				init:      nil,
			},
			{
				name:      "SQLite",
				dsnEnvKey: "RDBMS_SQLITE_DSN",
				init:      nil,
			},
			{
				name:      "InMemory",
				dsnEnvKey: "MEMORY_DSN",
				init:      nil,
			},
			{
				name:      "MongoDB",
				dsnEnvKey: "MONGODB_DSN",
				init:      nil,
			},
			{
				name:      "ElasticSearch",
				dsnEnvKey: "ELASTICSEARCH_DSN",
				init:      nil,
			},
		}
	)

	for _, s := range ss {
		t.Run(s.name, func(t *testing.T) {
			dsn, has := os.LookupEnv(s.dsnEnvKey)
			if !has {
				t.Skipf("no %s found, skipping %s persistence tests", s.dsnEnvKey, s.name)
				return
			}

			t.Logf("connecting to %s with %s", s.name, dsn)

			store, err := s.init(ctx, dsn)
			if err != nil {
				t.Errorf("failed to initialize %s storage", s.name)
				return
			}

			testAllGenerated(t, store)
		})
	}

}
