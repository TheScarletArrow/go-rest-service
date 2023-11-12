#FROM golang:alpine AS b
#WORKDIR /a
#ADD . /a
#RUN cd /a && go build -o go-r ./server.go
#
#FROM alpine
#WORKDIR /a
#COPY --from=b /a /a
#
#EXPOSE 8080
# Stage 1: Build Stage
FROM golang:alpine AS b
WORKDIR /a
ADD . /a
RUN cd /a && go build -o server.go

FROM alpine
WORKDIR /a
COPY --from=b /a /a

EXPOSE 8080
