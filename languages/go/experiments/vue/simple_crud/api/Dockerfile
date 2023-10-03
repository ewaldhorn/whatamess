###############
# BUILD STAGE #
###############
FROM golang:bookworm AS builder 

LABEL maintainer="Ewald Horn <ewald@nofuss.co.za>"

# Set the current working directory inside the container.
WORKDIR /app

# Copy dependency information to speed things up a bit.
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are 
# not changed between runs, in theory at least.
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container.
COPY . .

# Build the Go app.
RUN go build -o main *.go

#===================================================================================================

#############
# RUN STAGE #
#############

# Use the same base distribution as the build image.
FROM debian:bookworm

# We need CA certificates, so add them
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

# Set the current working directory inside the container.
WORKDIR /app

# Copy our executable from the build stage image into the current working directory.
COPY --from=builder /app/main .

# Now the environment variables
# Only used during debugging
# COPY app.env .

# And the Swagger docs
RUN mkdir docs
COPY docs docs

# Now the administrative console
RUN mkdir console
COPY console console

# Default landing page(s)
RUN mkdir www
COPY www www

# We also need the DB Migration scripts
RUN mkdir db
RUN mkdir db/migration
COPY db/migration db/migration

# Export Ports 80
EXPOSE 80

# Now create our user for running the app
RUN useradd -u 5000 apprunner
RUN chown -R apprunner /app
USER apprunner

CMD ["/app/main"]
