FROM ubuntu
RUN curl -L "https://cli.run.pivotal.io/stable?release=linux64-binary&version=v7" | tar -zx -C /usr/local/bin
