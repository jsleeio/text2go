FROM scratch
COPY text2go  /text2go
ENTRYPOINT ["/text2go"]
