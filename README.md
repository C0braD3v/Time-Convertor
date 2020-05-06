# Time Convertor
Convert timestamps to different formats, types, and timezones, such as UNIX, ISO, RFC1123, Local, or any IANA Timezone Database zone!

# Installation
## Windows
On Windows, Grab the newest [release](https://github.com/C0braD3v/Time-Convertor/releases) executable file, install it, and then run it the same as any other .exe. Once prompted, enter a timestamp in the specified format. Then, once prompted, enter a timestamp to convery to. Valid timestamps can be found [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

## Debian-Bassed Linux
Prerequisites: Must have git installed. Install using `sudo apt install git-all`.
On Debian, or any Debian-bassed Linux distro such as Ubuntu, run `sudo apt update`, then run `sudo apt install golang-go` to install GoLang. Then, clone this repository using `sudo git clone https://github.com/C0braD3v/Time-Convertor`. Once cloned, run `sudo go build convert.go` to build the binary, then run `./convert` to run the binary. (You could also run `sudo go run convert.go` if you do not wish to build the binary) Once prompted, enter a timestamp in the specified format. Then, once prompted, enter a timestamp to convery to. Valid timestamps can be found [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). Once the binary is built, feel free to remove GoLang using `sudo apt remove golang-go`, and then run `sudo apt-get autoremove`.

## MacOS
Prerequisites: Must have git installed. Install using `brew install git`.
On MacOS, run `brew update` to update HomeBrew, then run `brew install golang` to install GoLang. After this, clone this repository using `git clone https://github.com/C0braD3v/Time-Convertor`. Once cloned, run `go build convert.go` to build the binary, then run `./convert` to run the binary. (You could also run `go run convert.go` if you do not wish to build the binary) Once prompted, enter a timestamp in the specified format. Then, once prompted, enter a timestamp to convery to. Valid timestamps can be found [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). Once the binary is built, feel free to remove GoLang using `brew remove golang`.
