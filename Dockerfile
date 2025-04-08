# ---------- Etap 1: budowanie aplikacji ----------
    FROM golang:alpine AS builder

    ARG VERSION
    WORKDIR /app
    
    COPY main.go .
    
    RUN go mod init scratchapp && \
        go build -ldflags="-s -w -X main.version=${VERSION}" -o app
    
    # ---------- Etap 2: finalny obraz ----------
    FROM scratch
    
    COPY --from=builder /app/app /app
    
    EXPOSE 8080
    ENTRYPOINT ["/app"]
    