FROM golang:1.13-alpine3.10
WORKDIR /go-memo
RUN apk --no-cache git 
COPY . .
RUN go build
EXPOSE 3000
ENTRYPOINT ./go-memo

FROM node:latest
WORKDIR /go-memo
COPY web .                                                                                                                                               
RUN yarn && yarn build
EXPOSE 8080
ENTRYPOINT yarn start