# grape-vs-martini

API server example to compare Ruby's [Grape](https://github.com/intridea/grape) Web Framwork to [Martini](https://github.com/codegangsta/martini) Web Framwork in Go

[See blog post for details](http://stevenyue.com/2014/02/10/from-rubys-grape-to-martini-in-go-for-building-web-api-server/)

And would respond to the following routes:

    GET /projects(.json)
    GET /project/:id(.json)

Response format would be like:

	{"status": "Success", "data": [...]}
	{"status": "Fail", "error_message": "Bad api key"}

An example `curl` command would be like:
	
	curl "http://127.0.0.1:8080/projects?key=61c2339c1bc92bc48120b55513cd568b"

For Grape, `rack app` uses port `9292`, `rails app` uses port `3000` as default.

For Martini, port `8080` is used.

## How to run the server

### Grape Example

Under `grape-example`, there are two folders, in which Grape API is mounted on *Rack* or *Rails*, using [puma](https://github.com/puma/puma)

#### On Rack

1. Modify database settings in `config.ru` 

2. Run `bundle`

3. Run `rackup -o 127.0.0.1`

#### On Rails

1. Modify database settings in `config/database.yml` 

2. Run `bundle`

3. Run `rails s puma -b 127.0.0.1`

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

## Benchmark

A simple `ab` test is made, with command `ab -c 10 -n 1000 http://127.0.0.1:8080/projects?key=61c2339c1bc92bc48120b55513cd568b`

Test environment: Macbook Air CPU 1.7GHz Core i5, 8GB DDR3, OSX 10.9.1

ruby -v 2.0.0p247, go version go1.2 darwin/amd64, rails 3.2.16

### Grape

#### On Rack

	Concurrency Level:      10
	Time taken for tests:   30.643 seconds
	Complete requests:      1000
	Failed requests:        0
	Write errors:           0
	Total transferred:      2303000 bytes
	HTML transferred:       2211000 bytes
	Requests per second:    32.63 [#/sec] (mean)
	Time per request:       306.431 [ms] (mean)
	Time per request:       30.643 [ms] (mean, across all concurrent requests)
	Transfer rate:          73.39 [Kbytes/sec] received

	Connection Times (ms)
	              min  mean[+/-sd] median   max
	Connect:        0    0   0.1      0       1
	Processing:    55  306 505.3    267    5302
	Waiting:       52  290 505.0    256    5298
	Total:         55  306 505.3    267    5302

#### On Rails
	
	Concurrency Level:      10
	Time taken for tests:   3.781 seconds
	Complete requests:      1000
	Failed requests:        0
	Write errors:           0
	Non-2xx responses:      1000
	Total transferred:      199000 bytes
	HTML transferred:       9000 bytes
	Requests per second:    264.47 [#/sec] (mean)
	Time per request:       37.811 [ms] (mean)
	Time per request:       3.781 [ms] (mean, across all concurrent requests)
	Transfer rate:          51.40 [Kbytes/sec] received

	Connection Times (ms)
	              min  mean[+/-sd] median   max
	Connect:        0    0   0.2      0       3
	Processing:     3   37  64.4     28     718
	Waiting:        3   37  64.4     28     718
	Total:          3   38  64.5     28     719


### Martini

	Concurrency Level:      10
	Time taken for tests:   0.900 seconds
	Complete requests:      1000
	Failed requests:        0
	Write errors:           0
	Total transferred:      2314000 bytes
	HTML transferred:       2211000 bytes
	Requests per second:    1110.80 [#/sec] (mean)
	Time per request:       9.003 [ms] (mean)
	Time per request:       0.900 [ms] (mean, across all concurrent requests)
	Transfer rate:          2510.14 [Kbytes/sec] received

	Connection Times (ms)
	              min  mean[+/-sd] median   max
	Connect:        0    0   0.3      0       3
	Processing:     3    9   2.9      8      22
	Waiting:        3    8   2.9      8      22
	Total:          3    9   2.9      8      23
