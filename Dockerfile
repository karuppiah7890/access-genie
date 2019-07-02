FROM golang AS build
WORKDIR /app/access-genie
ADD . .
RUN make build

FROM alpine
COPY --from=build /app/access-genie/out/access-genie /usr/bin
RUN apk add ca-certificates
CMD ["access-genie"]