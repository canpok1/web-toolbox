# === ビルド用(バックエンド) ===
FROM golang:1.23-bookworm AS backend-builder
WORKDIR /app
COPY . .

RUN cd backend && make build

# === ビルド用(フロントエンド) ===
FROM node:22-bookworm AS frontend-builder
WORKDIR /app
COPY . .

RUN cd frontend && npm ci && npm run build

# === 実行用 ===
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=backend-builder /app/backend/build/release/server ./
COPY --from=frontend-builder /app/frontend/dist ./public

ENV TZ=Asia/Tokyo
ENV STATIC_DIR=/app/public
ENV REDIS_ADDRESS=redis:6379
ENV PORT=8080

EXPOSE 8080

ENTRYPOINT [ "/app/server" ]
