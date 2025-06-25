FROM golang:1.23 AS build

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS - HW15"
LABEL SERVICE="calendar"

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/calendar/config.yaml
COPY configs/config.yaml ${CONFIG_FILE}

EXPOSE 8083

CMD ${BIN_FILE} -config /etc/calendar