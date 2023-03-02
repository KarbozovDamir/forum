FROM golang:1.16

LABEL "site.name"="FORUM" \
      release-date="March, 2023" \
      description="forum" \
      authors="nrblzn | damirkap89"
WORKDIR /forum


# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy necessary files
ADD cmd cmd/
ADD configs configs/
ADD data data/
ADD internal internal/
ADD web web/

# Build application
RUN go build -o main cmd/main.go

EXPOSE 8181

CMD ["cmd/main.go"]
