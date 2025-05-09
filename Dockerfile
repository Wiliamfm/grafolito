FROM docker.io/golang:latest as base

FROM base as home

RUN useradd -m grafolito
USER grafolito
WORKDIR /home/grafolito/app

FROM home as build

COPY --chown=grafolito:grafolito go.mod go.sum ./
RUN go mod download

COPY --chown=grafolito:grafolito . .
RUN mkdir -p /home/grafolito/app/build
RUN go build -v -o /home/grafolito/app/build ./...

FROM home as release

USER grafolito
COPY --from=build /home/grafolito/app/build /usr/local/bin

CMD ["http"]
