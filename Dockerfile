FROM ubuntu:latest

LABEL maintainer "user"

ENV USER user
ENV HOME /home/${USER}
ENV SHELL /bin/bash

ARG HOST_ARCH="amd64"

RUN useradd -m ${USER}
RUN gpasswd -a ${USER} sudo
RUN echo "${USER}:work" | chpasswd
SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y
RUN apt-get install -y git sudo curl gcc sqlite3 make

RUN curl -O https://dl.google.com/go/go1.19.1.linux-${HOST_ARCH}.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-${HOST_ARCH}.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV PATH $PATH:${HOME}/go/bin
RUN go version

WORKDIR /home/${USER}

COPY . .

CMD ["make", "run"]