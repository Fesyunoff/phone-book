package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fesyunoff/phone-book/pkg/config"
	"github.com/fesyunoff/phone-book/pkg/service/dto"
	"github.com/fesyunoff/phone-book/pkg/storage"
	_ "github.com/lib/pq"
)

type PostgreBookStorage struct {
	Conn *sql.DB
	Conf *config.Config
}

var _ storage.Book = (*PostgreBookStorage)(nil)

func PreparePostgresDB(db *sql.DB, c *config.Config) {

	req := fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS  %s;`, c.SchemaName)
	result, err := db.Exec(req)
	CheckError(result, err, "Schema created")

	req = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(60) NOT NULL,
	"number" VARCHAR(16) NOT NULL UNIQUE,
	"note" VARCHAR(160)
	);`, c.SchemaName, c.TableName)
	result, err = db.Exec(req)
	CheckError(result, err, "Table created")
}

func (s *PostgreBookStorage) CreateEntry(e dto.Entry) (msg string, err error) {
	req := fmt.Sprintf(`INSERT INTO %s.%s(name, number, note) VALUES ('%s', '%s', '%s');`,
		s.Conf.SchemaName, s.Conf.TableName, e.Name, e.Number, e.Note)
	fmt.Println(req)
	_, err = s.Conn.Exec(req)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		msg = "CREATE"
	}

	return
}

func (s *PostgreBookStorage) DisplayEntry(e dto.Entry) (out []*dto.Entry, err error) {
	var req string
	if e.Name != "" {
		req = fmt.Sprintf("SELECT * FROM %s.%s WHERE name = '%s';", s.Conf.SchemaName, s.Conf.TableName, e.Name)
	} else if e.Number != "" {
		req = fmt.Sprintf("SELECT * FROM %s.%s WHERE number  = '%s';", s.Conf.SchemaName, s.Conf.TableName, e.Number)
	} else {
		err = errors.New("prepare request: uncorrect request format")
		return
	}
	row, err := s.Conn.Query(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer row.Close()
	for row.Next() {
		r := dto.Entry{}
		err = row.Scan(&r.Id, &r.Name, &r.Number, &r.Note)
		if err != nil {
			log.Fatalln(err.Error())
		}
		out = append(out, &r)
	}
	return
}

func (s *PostgreBookStorage) UpdateEntry(e dto.Entry) (msg string, err error) {
	req := fmt.Sprintf(`UPDATE %s.%s SET 
	name = '%s', number = '%s', note = '%s'  WHERE id  = %d;`,
		s.Conf.SchemaName, s.Conf.TableName,
		e.Name, e.Number, e.Note, e.Id)

	fmt.Println(req)
	_, err = s.Conn.Exec(req)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		msg = "UPDATE"
	}

	return
}
func (s *PostgreBookStorage) DeleteEntry(e dto.Entry) (msg string, err error) {
	req := fmt.Sprintf(`DELETE FROM %s.%s  
	WHERE id  = %d;`,
		s.Conf.SchemaName, s.Conf.TableName,
		e.Id)
	_, err = s.Conn.Exec(req)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		msg = "DELETE"
	}

	return
}

func CheckError(result sql.Result, err error, msg string) {
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		rows, _ := result.RowsAffected()
		log.Printf("%s. %d rows affected.", msg, rows)
	}
}

func NewPostgreBookStorage(conn *sql.DB, conf *config.Config) *PostgreBookStorage {
	return &PostgreBookStorage{
		Conn: conn,
		Conf: conf,
	}
}
