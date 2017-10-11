package main

// Doc documents behavior
const Doc = `
timez is a cli tool for converting time between timezones

it thinks your timezone is, from highest priority to lowest priority:
- the contents of ~/.timezrc
- the output of date +%z
- UTC

invocation:

# output current local and current UTC:
timez 

# output current PT:
timez PT

# what is my local time in pacific time:
timez 2017-10-11 19:05:00 in America/Los_Angeles
timez 2017-10-11 19:05:00 to US/Pacific
timez 2017-10-11 19:05:00 to Pacific
timez 2017-10-11 19:05:00 to PT

# what is this pacific time in my local:
timez PT 2017-10-11 19:05:00

# what is this unix seconds in UTC:
timez 1507702299755

# what is this unix nanoseconds in UTC:
timez 1507702299755000000000
`
