ARG PROJECTNAME=provisioner
ARG GOREPO=github.com/kfei/$PROJECTNAME
ARG GOVER=1.11

FROM tsaikd/node-yarn:latest AS web
ARG PROJECTNAME
COPY web /$PROJECTNAME/web
WORKDIR /$PROJECTNAME/web
RUN yarn install
RUN yarn run test:unit
RUN yarn run build

FROM golang:$GOVER AS backend
ARG PROJECTNAME
ARG GOREPO
RUN go get -v "github.com/tsaikd/gobuilder"
COPY . /go/src/$GOREPO
WORKDIR /go/src/$GOREPO
RUN gobuilder --all --check --test --debug

FROM alpine:3.8
ARG PROJECTNAME
ENV PROJECTNAME $PROJECNAME
ARG GOREPO
RUN apk add --no-cache curl
COPY --from=web /$PROJECTNAME/web/dist /$PROJECTNAME/web/dist/
COPY --from=backend /go/src/$GOREPO/$PROJECTNAME /$PROJECTNAME/
WORKDIR /$PROJECTNAME
CMD ["./$PROJECTNAME"]
