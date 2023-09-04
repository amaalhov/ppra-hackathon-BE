FROM golang:1.21.0 as builder
WORKDIR /code
COPY main.go .
COPY go.mod .
COPY go.sum .

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
# ARG SKAFFOLD_GO_GCFLAGS
# RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o /app .
RUN CGO_ENABLED=0 go build -trimpath -o /app .

FROM gcr.io/distroless/base
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single

COPY --from=builder /app /app

CMD ["./app"]
