FROM golang:latest

WORKDIR /go/src/github.com/oatmealraisin/openshift-listener
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
