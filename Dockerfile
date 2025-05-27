FROM golang:1.24-bookworm

RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

COPY . /go/src
WORKDIR /go/src

#RUN ls -l /go/src

RUN cp /go/src/libtraceable.so /usr/lib/

#RUN --mount=type=ssh go mod tidy
#RUN --mount=type=ssh make install-libs
RUN --mount=type=ssh CGO_ENABLED=1 TA_LOG_LEVEL=DEBUG go build -tags=traceable_filter -o /go/src/app.o .

ENV TA_LOG_LEVEL=DEBUG
ENTRYPOINT ["/go/src/app.o"]