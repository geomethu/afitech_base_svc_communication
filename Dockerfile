#build stage
FROM golang:1.17.8-alpine3.15 AS builder
RUN apk add --no-cache git
WORKDIR /go/src/github.com/geomethu/svc_comm

#Download necessary Go modules -> cache then then copy code so that we donot have to keep installing then when code changes
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /svc_comm ./cmd/main.go

#What command to run when our image is run as a container
CMD [ "/svc_comm" ]

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=svccommunication Version=0.0.1
# EXPOSE 3000
