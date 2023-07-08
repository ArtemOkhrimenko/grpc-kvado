package repo

type book struct {
	id   int    `db:"id"`
	name string `db:"name"`
}

type author struct {
	id   int    `db:"id"`
	name string `db:"name"`
}
