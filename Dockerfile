FROM debian:stable-slim

ENV PORT=8991

# copy source destination
COPY goserver /bin/goserver


CMD ["/bin/goserver"]
