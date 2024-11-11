# Unitags
MacOS proxy server for tracking Apple shit on other devices written in Golang

# How 2 use
## Prerequisites
- Golang
- MacOS <14.1

## Setting a Password
You're required to set a password for the webserver so people can't spy on your devices without credentials. The server will not run without it.
In `~/target/`, create a file called `password.txt` and put whatever passphrase you want for the webserver.

## Building & Running
```
go build -o target/unitags
cd target
./unitags
```
This will start a webserver at :4849

## Getting data from the server
To get an html file from the server, you must send a post request to `:4849/json` (yeah, I know the name doesn't make sense, but I have a plane to catch tomorrow) with the content of the body being the password you assigned in the `password.txt` file.

# Attribution
I nicked how the server gets the Find My data from another project called AirTagsAnywhere, which you can find [here](https://github.com/DylanAkp/AirtagsAnywhere). 

This project is licensed under the MIT license.

# Footnotes
I mostly just made this because I wanted to be able to keep my airtags after moving to Android, and I had an old MacBook lying around that wasn't doing anything, so I made this. I don't expect anybody to get any use out of this except for me, but in the rare case you find this to be a helpful project/tool/whatever, any contributions are appreciated.
