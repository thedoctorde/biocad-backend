# iron/go is the alpine image with only ca-certificates added
FROM iron/go
WORKDIR /app
# Now just add the binary
ADD myapp /app/
ENTRYPOINT ["./myapp"]