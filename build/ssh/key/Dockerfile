FROM ubuntu:latest

RUN apt-get update && apt-get install -y openssh-server

RUN mkdir /var/run/sshd

COPY ./dockerkey.pub /root/.ssh/authorized_keys

CMD ["/usr/sbin/sshd", "-D"]