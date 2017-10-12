timez 
=====

a cli tool for converting time between timezones
------------------------------------------------

### Usage:
```
timez <outputTZ...> <timestamp> <inputTZ>
```
NB: If inputTZ is not provided, timez currently assumes local timezone is Pacific/Auckland

### Done:
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
guess local timezone, from highest priority to lowest priority:
- the contents of ~/.timezrc
- the output of date +%z
- UTC

# handle numeric timezones
timez PT 2017-10-11 19:05:00 +0400

# what is this unix seconds in UTC:
timez 1507702299755

# what is this unix nanoseconds in UTC:
timez 1507702299755000000000
```
