package query

const (
	AddBook = `INSERT INTO db_go_sql.books
	(
		title,
		author,
		description
	)
	VALUES (
		$1, $2, $3
	)
	RETURNING *;`

	GetBooks = `SELECT * from db_go_sql.books ORDER BY id;`

	UpdateBook = `
	UPDATE db_go_sql.books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`

	DeleteBook = `
	DELETE from db_go_sql.books
	WHERE id = $1;
	`

	GetBookById = `
	SELECT * from db_go_sql.books 
	WHERE id = $1;`
)
