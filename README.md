timez
=====

a cli tool for converting time between timezones
------------------------------------------------

### Usage:
```
go get -u github.com/therealplato/timez
timez <outputTZ...> <timestamp> <inputTZ>
```
If `timestamp` is not provided, timez assumes now.

If outputTZ or inputTZ are not provided, timez uses the first available of:
- the contents of `~/.timezrc`
- the output of `date +%z`
- UTC


## Contributing:
`alias.go` contains a mapping from abbreviations I use to their authoritative zoneinfo string.

`formats.go` contains a slice of format strings. Parsing will try each format in order.

Feel free to PR new entries to these lists, what I've got so far is specific to my usecases.

## Examples:
```
# output current local and current UTC:
$ timez
Pacific/Auckland: 2017-10-12 22:30:27
UTC: 2017-10-12 09:30:27

# output current Pacific time twice and eastern time:
$ timez PT US/Pacific eastern
US/Pacific: 2017-10-12 02:32:13
US/Pacific: 2017-10-12 02:32:13
US/Eastern: 2017-10-12 05:32:13

# output current eastern, UTC:
$ timez ET UTC
US/Eastern: 2017-10-12 05:31:07
UTC: 2017-10-12 09:31:07

# given Pacific time, what is local time:
$ timez 2017-10-11 19:05:00 US/Pacific
Pacific/Auckland: 2017-10-12 15:05:00

# given local time, what is Pacific time:
$ timez US/Pacific 2017-10-11 19:05:00
US/Pacific: 2017-10-10 23:05:00

# given a eastern zoned time, what is pacific time:
$ timez PT 2017-10-11 19:05:00 ET
US/Pacific: 2017-10-11 16:05:00
```

### Todo:
```
# handle numeric timezones
timez PT 2017-10-11 19:05:00 +0400

# what is this unix seconds in UTC:
timez 1507702299755

# what is this unix nanoseconds in UTC:
timez 1507702299755000000000
```
