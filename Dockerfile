FROM alpine:latest

# Creates an app directory to hold your app’s source code
WORKDIR /

# Copies everything from your root directory into /app
ADD ./web ./web
ADD ./parSheet ./parSheet
COPY ./slot-server ./

# Tells Docker which network port your container listens on
EXPOSE 8081

RUN apk add --no-cache bash
# Specifies the executable command that runs when the container starts
CMD [ "/slot-server" ]