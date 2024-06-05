FROM golang:1.22.3 as base

FROM base as prod

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api

CMD ["air"]
