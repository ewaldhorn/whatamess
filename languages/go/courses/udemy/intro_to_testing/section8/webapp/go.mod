module webapp

go 1.22.0
toolchain go1.24.1

require github.com/go-chi/chi/v5 v5.0.12 // direct

require (
	github.com/alexedwards/scs/v2 v2.8.0
	github.com/jackc/pgconn v1.14.3
	github.com/jackc/pgx/v5 v5.7.1
	golang.org/x/crypto v0.35.0
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)
