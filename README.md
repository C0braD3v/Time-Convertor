# Time Convertor
This was a super simple GO script I wrote to convert a timestamp to different formats and types, such as UNIX, ISO, RFC1123, Local, or any IANA Timezone Databse zone. This program was nothing serious, just a something quick i made while bored to get timestamps in diffrent timezones and formats.

# Installation
If you're on Windows, simply download the EXE file, run it, and your command line will open. When prompted, enter in any timestamp in MM/dd/yyyy HH:mm:SS Z fortmat. Then, when prompted, Enter in any timezone in the IANA Timezone Database. The supported IANA timezones can be found here: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.

If you're on Linux, or MacOS, you will have to build to .go file yourself. To do this, set up a GO local enviormnet, clone the file into a working directory, and use `go build`.
