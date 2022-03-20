FROM golang:1.17-alpine AS gobuilder

LABEL maintainer="Julian Beck <mail@julianbeck.com (https://juli.sh/)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY adaptors/ ./adaptors
COPY messaging/ ./messaging
COPY storage/ ./storage
COPY utils/ ./utils
COPY main.go ./
COPY go.mod go.sum ./
# Set necessary environment variables needed for our image and build the API server.
RUN go build -ldflags="-s -w" -o pistock .

FROM node:13.12.0-alpine as nodebuilder

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

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=gobuilder ["/build/pistock", "/"]
COPY --from=nodebuilder ["/app/build/", "/frontend/build"] 
COPY website.yaml ./
# Command to run when starting the container.
ENTRYPOINT ["/pistock"]