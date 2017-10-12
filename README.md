timez 
=====

a cli tool for converting time between timezones
------------------------------------------------

### Usage:
```
go get -u github.com/therealplato/timez
timez <outputTZ...> <timestamp> <inputTZ>
```

If inputTZ is not provided, timez takes the first available of:
- the contents of `~/.timezrc`
- the output of `date +%z`
- UTC

### Details:
`alias.go` contains a mapping from abbreviations I use to their authoritative zoneinfo string. Feel free to PR new entries to this list.

```
# output current local and current UTC:
timez 

# output current Pacific time twice:
timez PT US/Pacific

# output current eastern, UTC:
timez ET UTC

# given Pacific time, what is local time :
timez 2017-10-11 19:05:00 US/Pacific

# what is this zoned time in pacific time:
timez PT 2017-10-11 19:05:00 ET 
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
