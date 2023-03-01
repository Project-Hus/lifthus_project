// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommunitiesColumns holds the columns for the "communities" table.
	CommunitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// CommunitiesTable holds the schema information for the "communities" table.
	CommunitiesTable = &schema.Table{
		Name:       "communities",
		Columns:    CommunitiesColumns,
		PrimaryKey: []*schema.Column{CommunitiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "community_id",
				Unique:  true,
				Columns: []*schema.Column{CommunitiesColumns[0]},
			},
		},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "uid", Type: field.TypeString},
		{Name: "revoked", Type: field.TypeBool, Default: false},
		{Name: "last_used_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "google_sub", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_verified", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true},
		{Name: "given_name", Type: field.TypeString},
		{Name: "family_name", Type: field.TypeString},
		{Name: "google_profile_picture", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommunitiesTable,
		RefreshTokensTable,
		UsersTable,
	}
)

func init() {
}
