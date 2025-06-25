FROM golang:1.24 AS build

ENV BIN_FILE /opt/calendar/calendar_sender
ENV CODE_DIR /go/src

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar_sender/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS - HW15"
LABEL SERVICE="calendar"

ENV BIN_FILE "/opt/calendar/calendar_sender"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/calendar/config-sender.yaml
COPY configs/config-sender.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config /etc/calendar