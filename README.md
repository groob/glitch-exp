## go-example

A very basic Glitch project showing how to get up and running with a Go web server. Note that Go is not an officially supported language on Glitch.

## Usage

Remix this instance and edit `server.go`. The majority of the magic for using Go is handled in the `package.json` startup script and `watch.json`.

Currently has 3 valid endpoints:

https://go-example.glitch.me/
 - Will return a hello message.
 
https://go-example.glitch.me/love/Go
 - Will return a message saying "I love Go" (or whatever the path is after /love/).
 
https://go-example.glitch.me/hacking
 - Will return a message about hacking away on Go.

## You are on your own

There are no tests, error checking, etc. Extending this is left as an excercise for the user.