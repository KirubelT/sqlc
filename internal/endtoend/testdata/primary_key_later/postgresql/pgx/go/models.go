// Code generated by sqlc. DO NOT EDIT.

package primary_key_later

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}