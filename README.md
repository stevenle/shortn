# URL shortener service.

Written in go. Runs on docker.

Note that this is just a toy I'm working on to play around with docker. It's not
intended for production services.

## Installation

```bash

# On Ubuntu 14.04:
git clone https://github.com/stevenle/shortn.git
sudo docker build -t stevenle/shortn shortn
sudo docker run -d -p 8080:8080 -p 9090:9090 stevenle/shortn shortn
```

## Example

```bash
# Make sure server's running:
http localhost:8080/_/ping  # should output "pong"

# Add a url:
http PUT localhost:9090/foo url=http://www.google.com

# Go to a url:
http GET localhost:8080/foo
```
