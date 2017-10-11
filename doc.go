package main

// Doc documents behavior
const Doc = `timez is a cli tool for converting time between timezones

most of this is still todo

 it thinks your timezone is, from highest priority to lowest priority:
- the contents of ~/.timezrc
- the output of date +%z
- UTC
- Pacific/Auckland

# output current local and current UTC:
timez 

# output current PT:
timez PT

# output current eastern, local, UTC:
timez ET @ UTC

# what is this local time in pacific time:
timez 2017-10-11 19:05:00 US/Pacific

# what is this zoned time in pacific time:
timez 2017-10-11 19:05:00 -0400 US/Pacific

# what is this pacific time in my local:
timez PT 2017-10-11 19:05:00

# what is this unix seconds in UTC:
timez 1507702299755

# what is this unix nanoseconds in UTC:
timez 1507702299755000000000
`
