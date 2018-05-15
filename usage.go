package main

const usage = `  timez <outputTZ...> <time> <inputTZ>

Omitting <outputTZ> or <inputTZ> assumes UTC.
Omitting <time> assumes now.
You can configure aliases in ~/.timezrc as yaml, e.g:
hiro: America/Los_Angeles
nell: Asia/Shanghai

https://github.com/therealplato/timez
`
