# grape-vs-martini

API server example to compare Ruby's [Grape](https://github.com/intridea/grape) Web Framwork to [Martini](https://github.com/codegangsta/martini) Web Framwork in Go

And would respond to the following routes:

    GET /projects(.json)
    GET /project/:id(.json)

Response format would be like:

	{"status": "Success", "data": [...]}
	{"status": "Fail", "error_message": "Bad api key"}

An example `curl` command would be like:
	
	curl "http://127.0.0.1:8080/projects?key=61c2339c1bc92bc48120b55513cd568b"

For Grape, port `9292` is used, and for Martini, port `8080` is used.

## How to run the server

### Grape Example

Under `grape-example` folder. The API is mounted directly on Rack, using [puma](https://github.com/puma/puma)

1. Modify database settings in `config.ru` 

2. Run `bundle`

3. Run `rackup`

### Martini Example

Under `martini-example` folder

1. Modify database settings in `server.go` file, change variables in `sql.Open()`

2. Run `go run server.go entry.go`

## How to create sample data

Example uses MySQL as default database.

Inside mysql, run:

	mysql> create database grape_vs_martini_api;

Dump sample data into database:

	mysql -u root -p grape_vs_martini_api < grape_vs_martini_api.sql

Sample data includes 10 companies, each with 50 projects, so there are 500 projects in total.

Sample companies' api keys:

	61c2339c1bc92bc48120b55513cd568b
	bb27bc6b7330aaac63ac809df83311b8
	e35402fa00728faf372abaf71623b7f4
	bca4bf9d131ab311760c5b790c8568a5
	49886a1c018496abb9ba4bf592c08d36
	850eedd94a4d9a01962628faac7ace91
	d036682c6ac91a41ac9eac064c02a43b
	77a767c39adc71fb241285309ff37ee5
	75e82d812630bf6d9371339f3e634801
	b5125eeabadcd5b3e55a76694d2b62b8