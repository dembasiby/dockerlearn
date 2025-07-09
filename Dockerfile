FROM debian:stable-slim

# copy source destination
COPY goserver /bin/goserver


CMD ["/bin/goserver"]
