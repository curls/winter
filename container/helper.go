package container

import "database/sql"

func DBFromContainer(key string) *sql.DB {

	item, err := Resolve(key)
	if err != nil {
		panic(err)
	}

	return item.(*sql.DB)
}
