# Hex Viewer CLI
Un outil en ligne de commande écrit en Go pour afficher le contenu d’un fichier sous forme d’hex dump, avec l’offset, les données en hexadécimal et leur équivalent ASCII.

## Fonctionnalités
- Affichage de l’offset en hexadécimal (8 chiffres).
- Lecture par blocs de taille configurable (par défaut 16 octets).
- Affichage des données en hexadécimal.
- Affichage du texte lisible (caractères ASCII imprimables), les autres sont remplacés par `.`.
- Gestion des fichiers binaires et textuels.

## Comment utiliser

1. **Compiler projet**
   ```go build .```
2. **Run the program**
   ./cli file_name
3. **Enter block size** (optional, default is 16):
- The program will ask for the block size (in bytes) at startup.

## Example output
```00000000 7061636b616765206d61696e0a package main.```
```00000010 0a696d706f727420280a202022 package main.```

## Project structure

- `main.go` : Program entry point.
- Function `PrintHexWithASCII` : Displays each block with offset, hex, and ASCII text.

## Requirements

- Go 1.19 or higher.
- A file to analyze.
