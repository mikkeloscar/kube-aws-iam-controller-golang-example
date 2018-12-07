FROM alpine

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates

ADD kube-aws-iam-controller-golang-example /example

ENTRYPOINT /example
