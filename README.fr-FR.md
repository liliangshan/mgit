# MGit - Outil de Gestion de Projets Git

## Introduction
MGit est un outil en ligne de commande pour gérer plusieurs projets Git. Il vous aide à gérer, synchroniser et mettre à jour efficacement plusieurs dépôts Git.

## Fonctionnalités Principales
- Gestion Multi-projets : Gérez plusieurs projets Git en un clic, améliorant l'efficacité du travail
- Support Multilingue : Prend en charge plusieurs langues dont le chinois (simplifié/traditionnel), l'anglais, le japonais, le coréen et le français
- Synchronisation de Base de Données Distante : Prend en charge la synchronisation des configurations de projet entre plusieurs appareils
- Contrôle de Version : Gestion de version intégrée avec mises à jour automatiques
- Intégration Système : Peut être ajouté aux variables d'environnement système pour une utilisation partout
- Gestion Intelligente des Branches : Détecte et change automatiquement les branches pour éviter les conflits
- Opérations par Lots : Prend en charge le pull/push simultané des modifications pour plusieurs projets

## Détails des Fonctionnalités

### 1. Collaboration Efficace
- Synchronisation automatique des configurations entre appareils pour une utilisation uniforme par l'équipe
- Gestion unifiée des projets d'équipe pour éviter les incohérences de configuration
- Les opérations par lots économisent du temps et améliorent l'efficacité
- Gestion intelligente des branches pour le changement et la synchronisation automatiques

### 2. Convivial
- Interface en ligne de commande interactive avec des opérations intuitives
- Support multilingue pour éliminer les barrières linguistiques
- Invites et aide intelligentes pour réduire la courbe d'apprentissage
- Format de commande unifié, facile à mémoriser et à utiliser

### 3. Sûr et Fiable
- Sauvegarde automatique des configurations pour prévenir les pertes accidentelles
- Protection par contrôle de version avec support des opérations de rollback
- Gestion intelligente des conflits pour assurer la sécurité des données
- Gestion des permissions pour prévenir les erreurs de manipulation

### 4. Hautement Extensible
- Prend en charge les configurations personnalisées pour répondre aux différents besoins
- Gestion flexible des projets adaptée à divers scénarios
- Mises à jour et améliorations continues pour une optimisation permanente
- Support des extensions plugin pour des fonctionnalités extensibles

## Notes de Version
1.0.16
- Ajout de la gestion d'instances d'application avec la commande `mgit new`
- Optimisation des règles de nommage des fichiers de configuration et de base de données
- Ajout de la configuration automatique des variables d'environnement
- Mise à jour forcée de la base de données distante (si activée)
- Ajout de la sélection de branche pour les pulls individuels
- Ajout de la fonctionnalité de changement de branches locales et distantes

## Nouvelles Fonctionnalités

### 1. Gestion des Instances d'Application
- Créer de nouvelles instances d'application :
  ```bash
  mgit new newapp  # Créer une nouvelle instance nommée newapp
  mgit new         # Création interactive
  ```
  - Les nouvelles applications héritent de toutes les fonctionnalités
  - Chaque instance possède ses propres fichiers de configuration et base de données
  - Les noms d'application sont automatiquement nettoyés des caractères invalides
  - Si l'application existe déjà, vous serez invité à utiliser l'existante

### 2. Règles de Nommage des Fichiers de Configuration
L'application sélectionne automatiquement les fichiers de configuration :
- mgit.exe : Utilise le fichier `.env`
- Autres applications : Utilise `.nomapp.env`
Exemple : gitmanager.exe -> .gitmanager.env

### 3. Règles de Nommage des Fichiers de Base de Données
Les fichiers de base de données suivent la même convention :
- mgit.exe : Utilise `projects.db`
- Autres applications : Utilise `nomapp.db`
Exemple : gitmanager.exe -> gitmanager.db

### 4. Répertoire de Base de Données Distante
Les répertoires de base de données distante utilisent également le nom de l'application :
- mgit.exe : Utilise le répertoire `.mgit_db`
- Autres applications : Utilise `.nomapp_db`
Exemple : gitmanager.exe -> .gitmanager_db

### 5. Configuration des Variables d'Environnement
Configuration automatique du répertoire de l'application dans les variables d'environnement :
```bash
mgit env  # Crée automatiquement MGIT_PATH ou nomapp_PATH
```
- Les noms des variables sont générés selon le nom de l'application
- Vérifie automatiquement l'existence de variables identiques
- Nécessite des droits administrateur sous Windows
- Certains paramètres peuvent nécessiter un redémarrage

### Notes Importantes
1. Lors de la création de nouvelles instances :
   - La nouvelle application démarre automatiquement
   - Les nouvelles applications ont des configurations et données indépendantes
   - Impossible de créer si une application du même nom existe

2. Configuration des Variables d'Environnement :
   - Nécessite des droits administrateur sous Windows
   - Peut nécessiter un redémarrage pour l'application des variables
   - Chemins en anglais recommandés pour éviter les problèmes d'encodage

3. Stockage des Données :
   - Tous les fichiers de configuration et bases de données sont stockés dans le répertoire de l'application
   - Le répertoire de base de données distante est géré comme un dépôt Git indépendant
   - Supporte la synchronisation automatique entre appareils

## Guide d'Installation

### Configuration Initiale
```bash
# Initialiser l'outil
./mgit init
# Ou
./mgit init mgit
```

### Configuration de l'Environnement
```bash
# Définir l'identifiant de la machine
./mgit set machine your-machine-name

# Définir le chemin de stockage des projets
./mgit set path /your/custom/path

# Voir la configuration actuelle
./mgit set
```

## Configuration de Base
L'outil utilise un fichier .env pour stocker la configuration :

```env
# Identifiant de la machine
MACHINE_ID=machine-01

# Chemin de l'application (répertoire parent pour tous les projets)
APP_PATH=/path/to/projects

# Configuration des variables d'environnement
MGIT_LANG=fr-FR
```

## Description de la Base de Données et des Fichiers de Configuration

### Règles de Nommage des Fichiers
L'application sélectionne automatiquement les fichiers selon le nom du programme :

1. Fichiers de Configuration :
   - mgit.exe : Utilise `.env`
   - Autres applications : Utilise `.nomapp.env`
   Exemple : gitmanager.exe -> .gitmanager.env

2. Base de Données Locale :
   - mgit.exe : Utilise `projects.db`
   - Autres applications : Utilise `nomapp.db`
   Exemple : gitmanager.exe -> gitmanager.db

3. Répertoire de Base de Données Distante :
   - mgit.exe : Utilise `.mgit_db`
   - Autres applications : Utilise `.nomapp_db`
   Exemple : gitmanager.exe -> .gitmanager_db

### Synchronisation de la Base de Données Distante
```bash
./mgit set
# Sélectionner "Activer la synchronisation du dépôt de base de données"
# Entrer l'adresse du dépôt de base de données
```

### Mécanisme de Synchronisation
- Synchronisation automatique après chaque push
- Récupération automatique de la dernière configuration avant chaque pull
- Support de la collaboration multi-utilisateurs avec fusion automatique

### Contenu Synchronisé
- Informations de configuration des projets (stockées dans nomapp.db)
- Paramètres des branches
- Derniers enregistrements de commit
- Identifiants des appareils

### Emplacements de Stockage
1. Base de Données Locale :
   - Stockée dans le répertoire de l'application
   - Fichiers de base de données créés automatiquement selon le nom de l'application

2. Base de Données Distante :
   - Stockée dans le répertoire .nomapp_db sous APP_PATH
   - Gérée comme un dépôt Git indépendant
   - Synchronisation et fusion automatiques entre appareils

### Notes Importantes
- Les fichiers de base de données sont créés automatiquement à la première utilisation
- La synchronisation distante nécessite une adresse de dépôt Git valide
- Les fichiers sont sauvegardés automatiquement pour prévenir les pertes
- Les données existantes sont migrées automatiquement lors du changement de base distante

## Guide d'Utilisation des Commandes

### Créer une Instance d'Application
```bash
# Créer une nouvelle instance (nom spécifié)
./mgit new project_manager

# Créer une nouvelle instance (interactif)
./mgit new
# Suivre les instructions pour entrer le nom
```

### Initialiser un Projet
```bash
# Initialiser un nouveau projet
./mgit init project_name https://github.com/user/repo.git

# Créer un nouveau projet
./mgit init mgit
# Ou simplement
./mgit init
```

### Tirer le Code (Pull)
```bash
# Tirer un projet unique
./mgit pull project_name
# Ou
./mgit pull
# Puis sélectionner le projet dans le menu

# Tirer tous les projets
./mgit pull-all
```

### Pousser le Code (Push)
```bash
# Pousser un projet unique
./mgit push project_name
# Ou
./mgit push
# Puis sélectionner le projet dans le menu
# Entrer le message de commit (vide pour utiliser le nom de la machine)

# Pousser tous les projets
./mgit push-all
```

### Gestion des Branches
```bash
# Configuration interactive des branches
./mgit branch

# Configurer les branches locale et distante d'un projet
./mgit branch project_name local_branch remote_branch
```

### Configurer la Branche de Pull
```bash
# Configuration interactive
./mgit set pull-branch

# Configurer pour un projet spécifique
./mgit set pull-branch project_name branch_name
```

### Gestion des Projets
```bash
# Voir la liste des projets
./mgit list

# Supprimer un projet
./mgit delete
# Sélectionner le projet à supprimer dans le menu
```

### Voir l'Aide
```bash
./mgit help
# Ou utiliser les alias
./mgit h
./mgit -h
./mgit -help
```

## Exemples de Menus Interactifs

### Menu de Sélection de Projet
1. Pull/Push d'un Projet Unique :
```bash
? Sélectionner le projet à opérer (Utiliser les flèches)
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [Annuler]
```

2. Gestion des Branches :
```bash
? Sélectionner le projet pour configurer les branches
❯ project1 (actuel : main -> origin/main)
  project2 (actuel : develop -> origin/develop)
  [Annuler]

? Entrer le nom de la branche locale : main
? Entrer le nom de la branche distante : origin/main
```

3. Suppression de Projet :
```bash
? Sélectionner le projet à supprimer
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [Annuler]

? Confirmer la suppression de project1 ? (O/n)
```

## Guide des Variables d'Environnement

### Utilisation de la Commande
```bash
mgit env
```

### Description des Fonctionnalités
- Configure automatiquement le répertoire de l'application comme variable d'environnement
- Le nom de la variable est généré selon le nom de l'application : `nomapp_PATH`
- Exemples :
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### Configuration des Variables d'Environnement Système

Ajouter MGit au PATH système pour un accès global :

1. Windows :
```powershell
# Ajouter le répertoire MGit au PATH utilisateur
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS :
```bash
# Éditer ~/.bashrc ou ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

Une fois la configuration terminée, vous pouvez utiliser la commande `mgit` depuis n'importe quel répertoire. L'outil gère automatiquement la gestion des projets, la synchronisation et les paramètres d'environnement selon votre configuration.

### Description des Commandes

#### Configuration du Proxy
```bash
mgit proxy
```
Configurer le serveur proxy Git, prenant en charge les proxys HTTP et HTTPS. Après l'exécution de la commande, le système vous demandera :
1. Si vous souhaitez utiliser un proxy
2. Sélectionner le type de proxy (HTTP/HTTPS)
3. Entrer l'adresse IP du serveur proxy
4. Entrer le port du serveur proxy

Exemple :
```bash
$ mgit proxy
Utiliser un proxy ? (o/N): o
Sélectionnez le type de proxy :
1) HTTP
2) HTTPS
Veuillez choisir : 1
Entrez l'IP du proxy : 127.0.0.1
Entrez le port du proxy : 7890
Configuration du proxy appliquée avec succès
```

```tool_call:save_file
path: README.fr-FR.md
content: [以上法文版本的完整内容]
```
