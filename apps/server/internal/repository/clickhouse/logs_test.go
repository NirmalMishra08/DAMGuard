package clickhouse

import (
	"context"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type fakeDriverConn struct {
	execCalled bool
	query      string
	args       []any
	execErr    error
}

func (f *fakeDriverConn) Contributors() []string { return nil }

func (f *fakeDriverConn) ServerVersion() (*driver.ServerVersion, error) {
	return nil, nil
}

func (f *fakeDriverConn) Select(context.Context, any, string, ...any) error {
	return nil
}

func (f *fakeDriverConn) Query(context.Context, string, ...any) (driver.Rows, error) {
	return nil, nil
}

func (f *fakeDriverConn) QueryRow(context.Context, string, ...any) driver.Row {
	return nil
}

func (f *fakeDriverConn) PrepareBatch(context.Context, string, ...driver.PrepareBatchOption) (driver.Batch, error) {
	return nil, nil
}

func (f *fakeDriverConn) Exec(ctx context.Context, query string, args ...any) error {
	f.execCalled = true
	f.query = query
	f.args = append([]any(nil), args...)
	return f.execErr
}

func (f *fakeDriverConn) AsyncInsert(context.Context, string, bool, ...any) error {
	return nil
}

func (f *fakeDriverConn) Ping(context.Context) error { return nil }

func (f *fakeDriverConn) Stats() driver.Stats { return driver.Stats{} }

func (f *fakeDriverConn) Close() error { return nil }

func TestInsertAuditEvent_UsesProvidedConnection(t *testing.T) {
	repoConn := &fakeDriverConn{}
	passedConn := &fakeDriverConn{}
	repo := &Repository{conn: repoConn}

	event := AuditEvent{
		EventID:      uuid.New(),
		Timestamp:    time.Unix(1710000000, 0).UTC(),
		DatabaseID:   "db-1",
		DatabaseName: "analytics",
		DatabaseType: "postgres",
		Username:     "alice",
		Query:        "SELECT 1",
		QueryType:    "read",
		ClientIP:     "127.0.0.1",
	}

	if err := repo.InsertQueryEvent(passedConn, event); err != nil {
		t.Fatalf("InsertAuditEvent returned error: %v", err)
	}

	if !passedConn.execCalled {
		t.Fatal("expected the provided connection to execute the insert")
	}

	if repoConn.execCalled {
		t.Fatal("expected the repository connection to remain unused")
	}

	if passedConn.query == "" {
		t.Fatal("expected the SQL insert statement to be built")
	}

	if len(passedConn.args) != 9 {
		t.Fatalf("expected 9 parameters, got %d", len(passedConn.args))
	}

	if got := passedConn.args[0].(uuid.UUID); got != event.EventID {
		t.Fatalf("expected first arg to be the event ID, got %v", got)
	}

	if got := passedConn.args[1].(time.Time); !got.Equal(event.Timestamp) {
		t.Fatalf("expected timestamp arg to match event timestamp, got %v", got)
	}
}
