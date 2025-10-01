FROM ubuntu

RUN apt-get update && apt-get install -y openssh-server

RUN mkdir /var/run/sshd

RUN echo "root:password" | chpasswd

RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config

RUN echo 'text in remote file' > remote.txt

CMD ["/usr/sbin/sshd", "-D"]