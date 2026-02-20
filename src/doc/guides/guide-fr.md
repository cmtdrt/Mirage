# Mirage — Guide utilisateur

Mirage est un outil en ligne de commande qui génère des API mock à partir d’une simple configuration JSON. Il permet de simuler des appels API réels, de prototyper des frontends et de tester des intégrations sans écrire de backend.

---

## Qu’est-ce que Mirage ?

Mirage lit un fichier JSON qui décrit des endpoints HTTP (méthode, chemin, corps de réponse, et optionnellement statut et délai). Il démarre un serveur HTTP local qui sert ces endpoints. N’importe quel client (navigateur, Postman, curl ou votre app) peut les appeler comme une vraie API.

**Cas d’usage :**
- Développement frontend avant que l’API réelle existe
- Tests d’intégration avec des réponses prévisibles
- Démonstrations et prototypes
- Apprentissage et expérimentation d’API HTTP

---

## Installation et lancement

### Compilation

```bash
go build -o mirage main.go
```

Ou avec le Makefile :

```bash
make build
```

### Démarrer le serveur

```bash
mirage serve <config.json>
```

Exemple :

```bash
mirage serve mirage.json
```

Le serveur écoute par défaut sur le **port 8080**.

---

## Options en ligne de commande

| Option | Description |
|--------|-------------|
| `mirage serve <config.json>` | Utiliser le fichier JSON donné comme configuration |
| `mirage serve` | Chercher `mirage.json` dans le répertoire courant ; erreur s’il est absent |
| `mirage serve --example` | Créer `mirage.example.json` à partir de l’exemple intégré et l’utiliser |
| `mirage serve --port=8081` | Démarrer sur le port 8081 (défaut : 8080) |
| `mirage guide-en` | Générer ce guide en anglais dans `mirage-guide-en.md` (puis quitter) |
| `mirage guide-fr` | Générer ce guide en français dans `mirage-guide-fr.md` (puis quitter) |

Les options peuvent être combinées, dans n’importe quel ordre :

```bash
mirage serve --example --port=3000
```

---

## Fichier de configuration

Le fichier de config est du JSON avec une seule clé racine : **`endpoints`**, un tableau d’objets endpoint.

### Champs obligatoires (par endpoint)

Ces champs doivent être présents pour chaque endpoint :

| Champ | Description |
|-------|-------------|
| `method` | Méthode HTTP : `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, etc. |
| `path` | Chemin URL (ex. `/api/users`). Variables de chemin supportées : `/users/{id}` |
| `response` | Corps de la réponse : chaîne, objet ou tableau (JSON) |

### Champs optionnels (par endpoint)

Ces champs peuvent être omis ; des valeurs par défaut ou aucun effet s’appliquent s’ils sont absents :

| Champ | Description |
|-------|-------------|
| `description` | Courte description affichée au démarrage du serveur |
| `status` | Code de statut HTTP (défaut : 200) |
| `delay` | Délai en **millisecondes** avant l’envoi de la réponse |

### Exemple minimal

```json
{
  "endpoints": [
    {
      "method": "GET",
      "path": "/hello",
      "response": "Bonjour le monde"
    }
  ]
}
```

Cela expose `GET /hello` et renvoie la chaîne `"Bonjour le monde"` avec le statut 200.

---

## Fonctionnalités en détail

### 1. Réponses statiques

Définir n’importe quel JSON (ou chaîne) comme `response`. Il est envoyé tel quel avec `Content-Type: application/json` (ou en chaîne).

```json
{
  "method": "GET",
  "path": "/api/config",
  "response": {
    "theme": "dark",
    "version": "1.0"
  }
}
```

### 2. Code de statut personnalisé

Utiliser `status` pour des codes de succès ou d’erreur (ex. 201, 404, 500) :

```json
{
  "method": "POST",
  "path": "/api/users",
  "status": 201,
  "response": { "id": 1, "username": "nouvel_utilisateur" }
}
```

### 3. Délai de réponse

Utiliser `delay` (en millisecondes) pour simuler un réseau ou un backend lent :

```json
{
  "method": "GET",
  "path": "/api/lent",
  "delay": 2000,
  "response": { "message": "Cela a pris 2 secondes" }
}
```

### 4. Variables de chemin

Les chemins peuvent contenir des segments dynamiques avec `{nom}`. Ils matchent n’importe quelle valeur pour ce segment.

Exemples :
- `/users/{id}` → matche `/users/1`, `/users/42`, `/users/abc`
- `/posts/{postId}/comments/{commentId}` → matche `/posts/10/comments/5`

```json
{
  "method": "GET",
  "path": "/api/v1/users/{id}",
  "response": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

La même réponse est renvoyée pour tout `id` ; les variables de chemin sont gérées par le routeur (Go 1.22+).

### 5. Descriptions

Utiliser `description` pour documenter les endpoints. Les descriptions sont affichées dans la console au démarrage du serveur :

```json
{
  "method": "GET",
  "path": "/health",
  "description": "Endpoint de santé",
  "response": { "status": "ok" }
}
```

---

## Démarrer rapidement

### 1. Lancer avec l’exemple intégré

```bash
mirage serve --example
```

Cela crée `mirage.example.json` et démarre le serveur avec. Pratique pour voir un exemple complet.

### 2. Lancer sur un port personnalisé

```bash
mirage serve mirage.json --port=3000
```

### 3. Utiliser le fichier de config par défaut

Si `mirage.json` existe dans le répertoire courant :

```bash
mirage serve
```

---

## Générer ce guide

Tu peux générer le guide utilisateur dans le répertoire courant :

- **Anglais :** `mirage serve --guide.en` → crée `mirage-guide-en.md`
- **Français :** `mirage serve --guide.fr` → crée `mirage-guide-fr.md`

Tu peux combiner avec d’autres options ; le guide est écrit avant le démarrage du serveur.

---

## En résumé

- **Config :** un fichier JSON avec un tableau `endpoints`.
- **Endpoints :** `method`, `path`, `response` ; optionnellement `description`, `status`, `delay`.
- **Chemins :** utiliser `{nomVariable}` pour les segments dynamiques.
- **CLI :** `serve`, `--example`, `--port=…`, et `guide-en` / `guide-fr` pour générer ce guide.

Pour plus d’exemples, exécuter `mirage serve --example` et ouvrir `mirage.example.json`.
