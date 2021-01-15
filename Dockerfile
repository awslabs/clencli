FROM ubuntu:latest AS ubuntu
WORKDIR /tmp
COPY build/clencli .
RUN ./clencli

FROM ubuntu:bionic AS bionic
WORKDIR /tmp
COPY build/clencli .
RUN ./clencli