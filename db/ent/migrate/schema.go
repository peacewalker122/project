// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountNotifsColumns holds the columns for the "account_notifs" table.
	AccountNotifsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeInt64},
		{Name: "notif_type", Type: field.TypeString, Size: 255},
		{Name: "notif_title", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "notif_content", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "notif_time", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AccountNotifsTable holds the schema information for the "account_notifs" table.
	AccountNotifsTable = &schema.Table{
		Name:       "account_notifs",
		Columns:    AccountNotifsColumns,
		PrimaryKey: []*schema.Column{AccountNotifsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "accountnotif_created_at",
				Unique:  false,
				Columns: []*schema.Column{AccountNotifsColumns[6]},
			},
		},
	}
	// NotifReadsColumns holds the columns for the "notif_reads" table.
	NotifReadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "notif_id", Type: field.TypeUUID, Unique: true},
		{Name: "account_id", Type: field.TypeInt64},
		{Name: "read_at", Type: field.TypeTime},
	}
	// NotifReadsTable holds the schema information for the "notif_reads" table.
	NotifReadsTable = &schema.Table{
		Name:       "notif_reads",
		Columns:    NotifReadsColumns,
		PrimaryKey: []*schema.Column{NotifReadsColumns[0]},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "token_type", Type: field.TypeString, Size: 255},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "raw", Type: field.TypeJSON},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountNotifsTable,
		NotifReadsTable,
		TokensTable,
	}
)

func init() {
}
