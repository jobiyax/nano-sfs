# nano sfs

Développé en Go 🐹, ce projet est un **serveur de fichiers statiques ultra léger ⚡**, permettant de servir rapidement vos fichiers depuis le dossier `static`. Pensé pour être simple, minimaliste et déployé partout facilement.

## Installation

```bash
# Cloner le dépôt
git clone https://github.com/jobiyax/nano-sfs.git

# Accéder au dossier
cd nano-sfs

# Nettoyer les dépendances
go mod tidy

# Lancer le serveur
go run main.go
```

## Utilisation

- Placez vos fichiers dans le dossier `static/`
- Ouvrez votre navigateur : `http://localhost:9000/`
- Vos fichiers seront accessibles directement depuis la racine
