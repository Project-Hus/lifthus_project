// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldGoogleSub holds the string denoting the google_sub field in the database.
	FieldGoogleSub = "google_sub"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailVerified holds the string denoting the email_verified field in the database.
	FieldEmailVerified = "email_verified"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGivenName holds the string denoting the given_name field in the database.
	FieldGivenName = "given_name"
	// FieldFamilyName holds the string denoting the family_name field in the database.
	FieldFamilyName = "family_name"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldProfilePictureURL holds the string denoting the profile_picture_url field in the database.
	FieldProfilePictureURL = "profile_picture_url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeHusSessions holds the string denoting the hus_sessions edge name in mutations.
	EdgeHusSessions = "hus_sessions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// HusSessionsTable is the table that holds the hus_sessions relation/edge.
	HusSessionsTable = "hus_sessions"
	// HusSessionsInverseTable is the table name for the HusSession entity.
	// It exists in this package in order to avoid circular dependency with the "hussession" package.
	HusSessionsInverseTable = "hus_sessions"
	// HusSessionsColumn is the table column denoting the hus_sessions relation/edge.
	HusSessionsColumn = "uid"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldProvider,
	FieldGoogleSub,
	FieldEmail,
	FieldEmailVerified,
	FieldName,
	FieldGivenName,
	FieldFamilyName,
	FieldBirthdate,
	FieldProfilePictureURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Provider defines the type for the "provider" enum field.
type Provider string

// Provider values.
const (
	ProviderHus    Provider = "hus"
	ProviderGoogle Provider = "google"
)

func (pr Provider) String() string {
	return string(pr)
}

// ProviderValidator is a validator for the "provider" field enum values. It is called by the builders before save.
func ProviderValidator(pr Provider) error {
	switch pr {
	case ProviderHus, ProviderGoogle:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for provider field: %q", pr)
	}
}
