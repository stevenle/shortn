# URL shortener service.

Written in go. Runs on docker.

```bash

git clone https://github.com/stevenle/shortn.git
sudo docker build -t stevenle/shortn .
sudo docker run -d -p 8080:8080 stevenle/shortn shortn
curl localhost:8080/ping  # should output "pong"
```
