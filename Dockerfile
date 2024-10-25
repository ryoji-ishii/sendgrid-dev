# Accept the Go version for the image to be set as a build argument.
ARG GO_VERSION=1.23.2

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -ldflags "-s -w" \
    -trimpath \
    -o /main ./main.go

# Final stage: the running container.
FROM gcr.io/distroless/static-debian12:nonroot AS final

# Import the compiled executable from the first stage.
COPY --from=builder --chown=nonroot:nonroot /main /main

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

WORKDIR /

# Perform any further action as an unprivileged user.
USER nonroot

# Run the compiled binary.
ENTRYPOINT ["/main"]
