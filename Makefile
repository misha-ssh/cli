# build and start ssh server with default port
# login - root
# address - localhost
# password - password
# port - 22
up-ssh:
	docker build -f ./build/ssh/default/Dockerfile -t ssh-host .
	docker run -d --name ssh-default -p 22:22 ssh-host

# stop and rm ssh-default container
down-ssh:
	docker stop ssh-default
	docker rm ssh-default

# build and start ssh server with 2222 port
# login - root
# address - localhost
# password - password
# port - 2222
up-ssh-port:
	docker build -f ./build/ssh/default/Dockerfile -t ssh-host .
	docker run -d --name ssh-port -p 2222:22 ssh-host

# stop and rm ssh-port container
down-ssh-port:
	docker stop ssh-port
	docker rm ssh-port

# generate ssh keys
# build and start ssh server with generated key
# login - root
# address - localhost
# private key - ./dockerkey
# port - 2222
up-ssh-key:
	ssh-keygen -b 4096 -t rsa -f dockerkey
	ssh-keygen -R localhost
	docker build -f ./build/ssh/key/Dockerfile -t ssh-host .
	docker run -d --name ssh-key -p 22:22 ssh-host

# rm ssh keys
# stop and rm ssh-key container
down-ssh-key:
	rm dockerkey dockerkey.pub
	docker stop ssh-key
	docker rm ssh-key

# use linter for formatted code
lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run