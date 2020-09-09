FROM scratch

ENV PORT 8080
EXPOSE $PORT
ADD apiproject /
ENTRYPOINT ["/apiproject","-config", "/etc/config.yml"]