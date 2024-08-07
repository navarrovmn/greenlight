This is the repo for the Greenlight application developed on Let's Go Further book, by Alex Edwards.

Dependencies:
* Greenlight DB-DSN: `export GREENLIGHT_DB_DSN='postgres://postgres:postgres@localhost/greenlight?sslmode=disable'`
* golang-migrate (`brew install golang-migrate`)
* `go run ./cmd/api`
* bcrypt

We are using MailTrap for mailing.

## History of Interesting Commands

* `migrate create -seq -ext=.sql -dir=./migrations create_movies_table`
* `migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up`
* [PostgreSQL Documentation](https://www.postgresql.org/docs/current/datatype.html)
* [Why to use Text instead of Char](https://www.depesz.com/2010/03/02/charx-vs-varcharx-vs-varchar-vs-text/)
