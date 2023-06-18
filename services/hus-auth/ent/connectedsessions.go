// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hus-auth/ent/connectedsessions"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ConnectedSessions is the model entity for the ConnectedSessions schema.
type ConnectedSessions struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hsid holds the value of the "hsid" field.
	Hsid uuid.UUID `json:"hsid,omitempty"`
	// Csid holds the value of the "csid" field.
	Csid uuid.UUID `json:"csid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ConnectedSessions) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connectedsessions.FieldID:
			values[i] = new(sql.NullInt64)
		case connectedsessions.FieldHsid, connectedsessions.FieldCsid:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ConnectedSessions", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ConnectedSessions fields.
func (cs *ConnectedSessions) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connectedsessions.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cs.ID = int(value.Int64)
		case connectedsessions.FieldHsid:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field hsid", values[i])
			} else if value != nil {
				cs.Hsid = *value
			}
		case connectedsessions.FieldCsid:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field csid", values[i])
			} else if value != nil {
				cs.Csid = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ConnectedSessions.
// Note that you need to call ConnectedSessions.Unwrap() before calling this method if this ConnectedSessions
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *ConnectedSessions) Update() *ConnectedSessionsUpdateOne {
	return NewConnectedSessionsClient(cs.config).UpdateOne(cs)
}

// Unwrap unwraps the ConnectedSessions entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *ConnectedSessions) Unwrap() *ConnectedSessions {
	_tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: ConnectedSessions is not a transactional entity")
	}
	cs.config.driver = _tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *ConnectedSessions) String() string {
	var builder strings.Builder
	builder.WriteString("ConnectedSessions(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cs.ID))
	builder.WriteString("hsid=")
	builder.WriteString(fmt.Sprintf("%v", cs.Hsid))
	builder.WriteString(", ")
	builder.WriteString("csid=")
	builder.WriteString(fmt.Sprintf("%v", cs.Csid))
	builder.WriteByte(')')
	return builder.String()
}

// ConnectedSessionsSlice is a parsable slice of ConnectedSessions.
type ConnectedSessionsSlice []*ConnectedSessions
