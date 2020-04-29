#######################################
# build stage for development 
#######################################
FROM golang:latest AS development

RUN apt-get update && apt-get install -y \
  git \
  build-essential

WORKDIR /go/src/app/server

COPY . .

RUN make tools-install
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

#######################################
# stage for production
#######################################
FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=development /go/src/app/server/app /usr/local/bin/
CMD ["app"]
