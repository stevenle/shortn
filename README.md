# URL shortener service.

Written in go. Runs on docker.

```bash

# On Ubuntu 14.04:
sudo apt-get update
sudo apt-get install git-core httpie lxc-docker

git clone https://github.com/stevenle/shortn.git
sudo docker build -t stevenle/shortn .
sudo docker run -d -p 8080:8080 stevenle/shortn shortn
http localhost:8080/ping  # should output "pong"

# Add a url:
http PUT localhost:9090/go/foo url=http://www.google.com

# Go to a url:
http GET localhost:8080/go/foo
```

Note that this is just a toy I'm working on to play around with docker. It's not
intended for production services.
