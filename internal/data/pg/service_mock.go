// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package pg

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"sync"
)

var (
	lockServiceMockConn            sync.RWMutex
	lockServiceMockDB              sync.RWMutex
	lockServiceMockExec            sync.RWMutex
	lockServiceMockGetSubordinates sync.RWMutex
	lockServiceMockPrepareQueries  sync.RWMutex
	lockServiceMockQuery           sync.RWMutex
	lockServiceMockSeed            sync.RWMutex
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             ConnFunc: func() *pgx.Conn {
// 	               panic("mock out the Conn method")
//             },
//             DBFunc: func() *pgx.Conn {
// 	               panic("mock out the DB method")
//             },
//             ExecFunc: func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
// 	               panic("mock out the Exec method")
//             },
//             GetSubordinatesFunc: func(ctx context.Context, userID string) ([]byte, error) {
// 	               panic("mock out the GetSubordinates method")
//             },
//             PrepareQueriesFunc: func(ctx context.Context, conn *pgx.Conn) error {
// 	               panic("mock out the PrepareQueries method")
//             },
//             QueryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
// 	               panic("mock out the Query method")
//             },
//             SeedFunc: func(ctx context.Context) error {
// 	               panic("mock out the Seed method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// ConnFunc mocks the Conn method.
	ConnFunc func() *pgx.Conn

	// DBFunc mocks the DB method.
	DBFunc func() *pgx.Conn

	// ExecFunc mocks the Exec method.
	ExecFunc func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	// GetSubordinatesFunc mocks the GetSubordinates method.
	GetSubordinatesFunc func(ctx context.Context, userID string) ([]byte, error)

	// PrepareQueriesFunc mocks the PrepareQueries method.
	PrepareQueriesFunc func(ctx context.Context, conn *pgx.Conn) error

	// QueryFunc mocks the Query method.
	QueryFunc func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// SeedFunc mocks the Seed method.
	SeedFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// Conn holds details about calls to the Conn method.
		Conn []struct {
		}
		// DB holds details about calls to the DB method.
		DB []struct {
		}
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SQL is the sql argument value.
			SQL string
			// Arguments is the arguments argument value.
			Arguments []interface{}
		}
		// GetSubordinates holds details about calls to the GetSubordinates method.
		GetSubordinates []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID string
		}
		// PrepareQueries holds details about calls to the PrepareQueries method.
		PrepareQueries []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Conn is the conn argument value.
			Conn *pgx.Conn
		}
		// Query holds details about calls to the Query method.
		Query []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SQL is the sql argument value.
			SQL string
			// Args is the args argument value.
			Args []interface{}
		}
		// Seed holds details about calls to the Seed method.
		Seed []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
}

// Conn calls ConnFunc.
func (mock *ServiceMock) Conn() *pgx.Conn {
	if mock.ConnFunc == nil {
		panic("ServiceMock.ConnFunc: method is nil but Service.Conn was just called")
	}
	callInfo := struct {
	}{}
	lockServiceMockConn.Lock()
	mock.calls.Conn = append(mock.calls.Conn, callInfo)
	lockServiceMockConn.Unlock()
	return mock.ConnFunc()
}

// ConnCalls gets all the calls that were made to Conn.
// Check the length with:
//     len(mockedService.ConnCalls())
func (mock *ServiceMock) ConnCalls() []struct {
} {
	var calls []struct {
	}
	lockServiceMockConn.RLock()
	calls = mock.calls.Conn
	lockServiceMockConn.RUnlock()
	return calls
}

// DB calls DBFunc.
func (mock *ServiceMock) DB() *pgx.Conn {
	if mock.DBFunc == nil {
		panic("ServiceMock.DBFunc: method is nil but Service.DB was just called")
	}
	callInfo := struct {
	}{}
	lockServiceMockDB.Lock()
	mock.calls.DB = append(mock.calls.DB, callInfo)
	lockServiceMockDB.Unlock()
	return mock.DBFunc()
}

// DBCalls gets all the calls that were made to DB.
// Check the length with:
//     len(mockedService.DBCalls())
func (mock *ServiceMock) DBCalls() []struct {
} {
	var calls []struct {
	}
	lockServiceMockDB.RLock()
	calls = mock.calls.DB
	lockServiceMockDB.RUnlock()
	return calls
}

// Exec calls ExecFunc.
func (mock *ServiceMock) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	if mock.ExecFunc == nil {
		panic("ServiceMock.ExecFunc: method is nil but Service.Exec was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		SQL       string
		Arguments []interface{}
	}{
		Ctx:       ctx,
		SQL:       sql,
		Arguments: arguments,
	}
	lockServiceMockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	lockServiceMockExec.Unlock()
	return mock.ExecFunc(ctx, sql, arguments...)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedService.ExecCalls())
func (mock *ServiceMock) ExecCalls() []struct {
	Ctx       context.Context
	SQL       string
	Arguments []interface{}
} {
	var calls []struct {
		Ctx       context.Context
		SQL       string
		Arguments []interface{}
	}
	lockServiceMockExec.RLock()
	calls = mock.calls.Exec
	lockServiceMockExec.RUnlock()
	return calls
}

// GetSubordinates calls GetSubordinatesFunc.
func (mock *ServiceMock) GetSubordinates(ctx context.Context, userID string) ([]byte, error) {
	if mock.GetSubordinatesFunc == nil {
		panic("ServiceMock.GetSubordinatesFunc: method is nil but Service.GetSubordinates was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID string
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	lockServiceMockGetSubordinates.Lock()
	mock.calls.GetSubordinates = append(mock.calls.GetSubordinates, callInfo)
	lockServiceMockGetSubordinates.Unlock()
	return mock.GetSubordinatesFunc(ctx, userID)
}

// GetSubordinatesCalls gets all the calls that were made to GetSubordinates.
// Check the length with:
//     len(mockedService.GetSubordinatesCalls())
func (mock *ServiceMock) GetSubordinatesCalls() []struct {
	Ctx    context.Context
	UserID string
} {
	var calls []struct {
		Ctx    context.Context
		UserID string
	}
	lockServiceMockGetSubordinates.RLock()
	calls = mock.calls.GetSubordinates
	lockServiceMockGetSubordinates.RUnlock()
	return calls
}

// PrepareQueries calls PrepareQueriesFunc.
func (mock *ServiceMock) PrepareQueries(ctx context.Context, conn *pgx.Conn) error {
	if mock.PrepareQueriesFunc == nil {
		panic("ServiceMock.PrepareQueriesFunc: method is nil but Service.PrepareQueries was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Conn *pgx.Conn
	}{
		Ctx:  ctx,
		Conn: conn,
	}
	lockServiceMockPrepareQueries.Lock()
	mock.calls.PrepareQueries = append(mock.calls.PrepareQueries, callInfo)
	lockServiceMockPrepareQueries.Unlock()
	return mock.PrepareQueriesFunc(ctx, conn)
}

// PrepareQueriesCalls gets all the calls that were made to PrepareQueries.
// Check the length with:
//     len(mockedService.PrepareQueriesCalls())
func (mock *ServiceMock) PrepareQueriesCalls() []struct {
	Ctx  context.Context
	Conn *pgx.Conn
} {
	var calls []struct {
		Ctx  context.Context
		Conn *pgx.Conn
	}
	lockServiceMockPrepareQueries.RLock()
	calls = mock.calls.PrepareQueries
	lockServiceMockPrepareQueries.RUnlock()
	return calls
}

// Query calls QueryFunc.
func (mock *ServiceMock) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if mock.QueryFunc == nil {
		panic("ServiceMock.QueryFunc: method is nil but Service.Query was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		SQL  string
		Args []interface{}
	}{
		Ctx:  ctx,
		SQL:  sql,
		Args: args,
	}
	lockServiceMockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	lockServiceMockQuery.Unlock()
	return mock.QueryFunc(ctx, sql, args...)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//     len(mockedService.QueryCalls())
func (mock *ServiceMock) QueryCalls() []struct {
	Ctx  context.Context
	SQL  string
	Args []interface{}
} {
	var calls []struct {
		Ctx  context.Context
		SQL  string
		Args []interface{}
	}
	lockServiceMockQuery.RLock()
	calls = mock.calls.Query
	lockServiceMockQuery.RUnlock()
	return calls
}

// Seed calls SeedFunc.
func (mock *ServiceMock) Seed(ctx context.Context) error {
	if mock.SeedFunc == nil {
		panic("ServiceMock.SeedFunc: method is nil but Service.Seed was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockServiceMockSeed.Lock()
	mock.calls.Seed = append(mock.calls.Seed, callInfo)
	lockServiceMockSeed.Unlock()
	return mock.SeedFunc(ctx)
}

// SeedCalls gets all the calls that were made to Seed.
// Check the length with:
//     len(mockedService.SeedCalls())
func (mock *ServiceMock) SeedCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockServiceMockSeed.RLock()
	calls = mock.calls.Seed
	lockServiceMockSeed.RUnlock()
	return calls
}
