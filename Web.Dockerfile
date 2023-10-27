FROM alpine:latest

# Creates an app directory to hold your app’s source code
WORKDIR /

ADD ./web ./web
COPY ./build/web-server ./

# Tells Docker which network port your container listens on
EXPOSE 8081

RUN apk add --no-cache bash
# Specifies the executable command that runs when the container starts
CMD [ "/web-server" ]