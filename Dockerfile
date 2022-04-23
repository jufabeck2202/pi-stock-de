FROM golang:1.17-alpine AS gobuilder

LABEL maintainer="Julian Beck <mail@julianbeck.com (https://juli.sh/)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY internal/ ./internal
COPY main.go ./
COPY go.mod go.sum ./
# Set necessary environment variables needed for our image and build the API server.
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
	go build -ldflags="-s -w" -o pistock .

FROM node:17.7.1-alpine as nodebuilder

# set working directory
WORKDIR /app
# install dependencies
COPY frontend/package.json frontend/tsconfig.json ./
COPY frontend/package-lock.json ./
RUN npm install --silent

# copy the code into the container
COPY frontend/src ./src
COPY frontend/public ./public

RUN npm run build

FROM alpine:latest

# Move Files from build steps into the container
COPY --from=gobuilder ["/build/pistock", "/"]
COPY --from=nodebuilder ["/app/build/", "/frontend/build"] 
COPY website.yaml ./
ENTRYPOINT ["/pistock"]