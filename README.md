# Hex Viewer CLI
Un outil en ligne de commande écrit en Go pour afficher le contenu d’un fichier sous forme d’hex dump, avec l’offset, les données en hexadécimal et leur équivalent ASCII.

## Fonctionnalités
- Affichage de l’offset en hexadécimal (8 chiffres).
- Lecture par blocs de taille configurable (par défaut 16 octets).
- Affichage des données en hexadécimal.
- Affichage du texte lisible (caractères ASCII imprimables), les autres sont remplacés par `.`.
- Gestion des fichiers binaires et textuels.
