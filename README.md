# freeCodeCamp-cheer

Desktop notifications when a freeCodeCamp.org user gets more points.

## Purpose

I have a friend who is starting to learn to how to program. They're using freeCodeCamp as their primary source for content right now. They sent me a link to their profile so I can checkout their work and encourage them to keep going. Amazed at how fast my friend was progressing through the curriculum I decided to make a little program to monitor the number of points on the profile and send me a little notification when they make progress. That way I can cheer them on!

## Usage

### Go version

I've been wanting to learn more go and thought this would be a good chance.

`go run fcc-cheer.go -user raybb -sec 10`

### Bash version

I first started this project with a simple, and very fragile, Bash script. It works and I wanted to keep it around in case I wanted to do similar scraping in the future.

`./fcc-cheer.sh raybb`

### Flags

* `-user` specify the username of the account you want to monitor
* `-sec` number of seconds between checks (default 10)
