# ESGO

## Development

1. Provision elastic search stack with Docker Compose.
```Bash
docker compose up
```
2. Copy HTTPS certficate from elastic search container for backend to use.
```
docker cp esgo-es01-1:/usr/share/elasticsearch/config/certs/ca/ca.crt .
```
