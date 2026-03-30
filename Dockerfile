# Étape 1 : Build avec Go 1.26
FROM golang:1.26-alpine AS builder

# Active la compilation statique
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copier les fichiers de module pour gérer les dépendances
COPY go.mod go.* ./
RUN go mod tidy

# Copier le code source
COPY . .

# Compiler le binaire sans debug pour réduire la taille
RUN go build -ldflags="-s -w" -o nano-sfs main.go

# Étape 2 : Image finale légère
FROM alpine:latest

WORKDIR /app

# Copier le binaire compilé depuis l’étape builder
COPY --from=builder /app/nano-sfs .

# Copier le dossier static
COPY --from=builder /app/static ./static

# Exposer le port utilisé
EXPOSE 9000

# Commande de démarrage
CMD ["./nano-sfs"]
