FROM golang:1.16

LABEL "site.name"="FORUM" \
      release-date="march, 2023" \
      description="forum" \
      authors="nrblzn and damirkap89"
WORKDIR /forum


# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy necessary files
ADD cmd cmd/
ADD configs configs/
ADD internal internal/
ADD web web/

# Build application
RUN go build -o main cmd/main.go

EXPOSE 8181

CMD ["cmd/main.go"]