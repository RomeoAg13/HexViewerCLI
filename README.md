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
2. **Executer**
   ./cli nom_du_fichier
3. **Entrer la taille du bloc** (optionnel, par défaut 16) :
- Le programme demande la taille du bloc (en octets) au démarrage.

## Exemple de sortie
```00000000 7061636b616765206d61696e0a package main.```
```00000010 0a696d706f727420280a202022 package main.```

## Structure

- `main.go` : Point d’entrée du programme.
- Fonction `PrintHexWithASCII` : Affiche chaque bloc avec offset, hex et texte.

## Prérequis

- Go 1.19 ou supérieur.
- Un fichier à analyser.

## Prochaines étapes

- Ajouter des options en ligne de commande (ex: `-c 16` pour la taille du bloc).
- Support de l’entrée standard (`-`).
- Affichage aligné et formaté.
- Gestion des erreurs plus robuste.

## Licence

MIT
