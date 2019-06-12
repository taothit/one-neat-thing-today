# one-neat-thing-today [![Build Status](https://travis-ci.org/taothit/one-neat-thing-today.svg?branch=master)](https://travis-ci.org/taothit/one-neat-thing-today)
Daily provider for an interesting concept, term, event, or idea.
_____
## Command-line Tool

### Usage
```shell
$ ontt <options>
[Neat Thing] - [Definition|Description]
(link to top example of thing)
```

### Options
|Short |Long|Parameters|Description|
|---|---|---|---|
|-a|--after-date|[date]|shows all neat things starting (inclusive) on date until $today (inclusive) |
|-b|--bibliography| |shows wikipedia-style bibliography with links|
|-e|--export|[start date] ([end date])| writes out formatted file of neat things starting at date until $today, or [end date] (if provided).|
|-f|format|[output format] ([formatter])|allows user to specify whether the response is structured or plaintext.  Initially, ONTT will provide output as plaintext, json, or xml. If [formatter] is provided, then ONTT will attempt to use the go command by piping the output as json into the  Formatter command.|
|-l|link| |limit output to only the link of the neat thing|
|-n|name| | limit output to only the name of the neat thing|
|-o|--output|[file path]|Writes neat thing(s) to out to [file path]
|-p|prior-days| [number of days]| shows [number of days] neat thing(s) before today|
|-t|timeline| [start date] ([end date])|output html page that shows horizontal timeline with snapshots of webpages when hovering of the date/thing; navigation tools will be displayed on the timeline when it spans a sufficient range of dates.|
|-u|--url|[host:port]|Directs ontt client to get new things from ontt server at [host:port]| 
|-v|version| | displays version of ontt command as well as api version of ontt service|
|-w|week| |shows the neat things from start of week until $today|
|-y|yesterday| |shows yesterday’s neat thing|
