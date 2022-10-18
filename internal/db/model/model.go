//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type ColumnsPerson struct {
	ID, Login, Password, CreatedAt, DeletedAt string
}

type ColumnsSt struct {
	Person ColumnsPerson
}

var Columns = ColumnsSt{
	Person: ColumnsPerson{
		ID:        "id",
		Login:     "login",
		Password:  "password",
		CreatedAt: "created_at",
		DeletedAt: "deleted_at",
	},
}

type TablePerson struct {
	Name, Alias string
}

type TablesSt struct {
	Person TablePerson
}

var Tables = TablesSt{
	Person: TablePerson{
		Name:  "person",
		Alias: "t",
	},
}

type Person struct {
	tableName struct{} `pg:"person,alias:t,discard_unknown_columns"`

	ID        string     `pg:"id,pk,type:uuid"`
	Login     string     `pg:"login,use_zero"`
	Password  []byte     `pg:"password"`
	CreatedAt *time.Time `pg:"created_at"`
	DeletedAt *time.Time `pg:"deleted_at,,soft_delete"`
}
