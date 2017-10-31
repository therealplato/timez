timez
=====

a cli tool for converting time between timezones
------------------------------------------------

## Usage:
```
go get -u github.com/therealplato/timez
timez <outputTZ...> <timestamp> <inputTZ>
```
If `timestamp` is not provided, timez assumes now.

If `outputTZ` or `inputTZ` are not provided, timez uses the first available of:
- a configured `default` alias in `~/.timezrc`
- the output of `date +%z`
- UTC

## Configuration:
See [`example.timezrc`](https://github.com/therealplato/timez/blob/master/example.timezrc)

## Examples:
```
# output current local and current UTC:
$ timez
Pacific/Auckland: 2017-10-12 22:30:27
UTC: 2017-10-12 09:30:27

# output current Pacific time twice and eastern time:
$ timez PT US/Pacific eastern
PT: 2017-10-12 02:32:13
US/Pacific: 2017-10-12 02:32:13
US/Eastern: 2017-10-12 05:32:13

# given Pacific time, what is local time:
$ timez 2017-10-11 19:05:00 US/Pacific
Pacific/Auckland: 2017-10-12 15:05:00

# given local time, what is Pacific time:
$ timez US/Pacific 2017-10-11 19:05:00
US/Pacific: 2017-10-10 23:05:00

# given a eastern zoned time, what is pacific time:
$ timez PT 2017-10-11 19:05:00 ET
US/Pacific: 2017-10-11 16:05:00

# given the aliases in example.timezrc are configured,
# what's nell's time when it's midnight for hiro?
$ timez nell 2017-11-01 00:00:00 hiro
nell: 2017-11-01 15:00:00
```

## Todo:
```
# handle numeric timezones
timez PT 2017-10-11 19:05:00 +0400

# what is this unix seconds in UTC:
timez 1507702299755

# what is this unix nanoseconds in UTC:
timez 1507702299755000000000
```

## Contributing:
`defaults.go` contains some timezone aliases and format strings that I personally use frequently.  
The aliases are mappings from abbreviations like "auckland, ET" to zoneinfo strings "Pacific/Auckland", "US/Eastern".  
Feel free to PR new entries to these lists. I'd also welcome a PR for configurable format strings.  
