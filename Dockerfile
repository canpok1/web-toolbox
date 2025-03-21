# === ビルド用 ===
FROM golang:1.23-bookworm AS backend-builder
WORKDIR /app
COPY . .

RUN cd backend && make build

# === 実行用 ===
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=backend-builder /app/backend/build/release/server ./

ENV TZ=Asia/Tokyo

EXPOSE 8080

ENTRYPOINT [ "/app/server" ]
