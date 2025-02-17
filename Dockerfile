# stage 1 : build app
FROM golang:1.23.6-alpine AS builder
# create and set working directory
WORKDIR /app
# copy package files
COPY go.mod go.sum ./
# download all the package
RUN go mod download
# copy the rest of source code
COPY . .  
# build the app
RUN go build -o /bin/main cmd/app/main.go

# stage 2 : run app
FROM alpine:latest
WORKDIR /app
# copy binary from the builder stage
COPY --from=builder /bin/main /bin/main
# copy .env file
COPY .env.example .env.example
# expose app port based on .env
EXPOSE 8081
# run the app
CMD ["/bin/main"]
