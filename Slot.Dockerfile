FROM alpine:latest

# Creates an app directory to hold your app’s source code
WORKDIR /

ADD ./parSheet ./parSheet
COPY ./build/slot-server ./

# Tells Docker which network port your container listens on
EXPOSE 8088

RUN apk add --no-cache bash
# Specifies the executable command that runs when the container starts
CMD [ "/slot-server" ]