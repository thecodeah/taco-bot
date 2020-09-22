case "$1" in
    "dev") docker-compose up;;
    "prod") docker-compose -f docker-compose.yml -f docker-compose.prod.yml -d up;;
    *) echo "Usage: $0 dev || prod"
esac