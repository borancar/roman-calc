FROM golang as build
WORKDIR /go/src/roman-calc

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/roman-calc

RUN dep ensure
RUN go build

FROM golang
WORKDIR /go/src/roman-calc
ENTRYPOINT ./roman-calc

COPY --from=build /go/src/roman-calc/roman-calc .
COPY --from=build /go/src/roman-calc/templates ./templates
COPY --from=build /go/src/roman-calc/swagger.yaml .
