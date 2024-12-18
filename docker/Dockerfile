FROM golang:alpine AS builder

LABEL stage=gobuilder

ARG BUILD_MODE

ENV CGO_ENABLED=0

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .
RUN if [[ "${BUILD_MODE}" == "debug" ]] ; then swag i -g cmd/main_debug.go -o docs/ --ot go ; fi
RUN go build --tags ${BUILD_MODE} --ldflags="-s -w" -buildvcs=false -o /app/calc ./cmd/


FROM scratch

WORKDIR /app
COPY --from=builder /app/calc /app/calc

CMD [ "./calc" ]