
.PHONY: dev dev-build dev-down dev-logs dev-clean help

dev:
	@echo "🚀 Démarrage de l'environnement de développement..."
	docker compose -f docker-compose.dev.yml up

dev-build:
	@echo "🔨 Construction et démarrage..."
	docker compose -f docker-compose.dev.yml up --build

dev-down:
	@echo "🛑 Arrêt de l'environnement de développement..."
	docker compose -f docker-compose.dev.yml down

dev-logs:
	@echo "📋 Logs de l'application..."
	docker compose -f docker-compose.dev.yml logs -f app

dev-clean:
	@echo "🧹 Nettoyage complet..."
	docker compose -f docker-compose.dev.yml down -v --remove-orphans
	docker system prune -f

dev-shell:
	@echo "🐚 Shell dans le conteneur de l'app..."
	docker compose -f docker-compose.dev.yml exec app sh

dev-db:
	@echo "🗄️ Connexion à PostgreSQL..."
	docker compose -f docker-compose.dev.yml exec postgres psql -U quillstream -d quillstream_dev

install-air:
	@echo "📦 Installation d'Air..."
	go install github.com/air-verse/air@latest

help:
	@echo "📚 Commandes disponibles:"
	@echo "  make dev        - Démarrer l'environnement de développement"
	@echo "  make dev-build  - Construire et démarrer"
	@echo "  make dev-down   - Arrêter l'environnement"
	@echo "  make dev-logs   - Voir les logs de l'app"
	@echo "  make dev-clean  - Nettoyage complet (⚠️ supprime les données)"
	@echo "  make dev-shell  - Shell dans le conteneur de l'app"
	@echo "  make dev-db     - Connexion à PostgreSQL"
	@echo "  make help       - Afficher cette aide"