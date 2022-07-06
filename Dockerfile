# Build container
FROM alpine:latest as BUILD

WORKDIR /build

# Set envs
ENV GO111MODULE=on

# Install dependencies 
RUN apk add --no-cache nodejs npm go make

# copy files to workdir
COPY . ./

# Install onkyo-remote dependencies
RUN make deps

# Compile the app
RUN make linux

# ==============================================================================
# App container
FROM alpine:latest 

RUN apk add --no-cache tini

WORKDIR /app

# Configuration files are stored here
VOLUME ["/config"]

# Copy executable to /app, and entrypoint to root (/)
COPY --from=BUILD /build/onkyo-remote-linux ./onkyo-remote
COPY --from=BUILD /build/entrypoint.sh /

# Run entrypoint.sh
ENTRYPOINT [ "tini", "--" ]
CMD ["/entrypoint.sh"]

# App runs on port 8080 by default (unless config is changed)
EXPOSE 8080/tcp
