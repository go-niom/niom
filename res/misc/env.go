package misc

const MiscEnv = `
# APP settings:
APP_NAME={{ .NameLowerCase}}
APP_HOST=127.0.0.1
APP_PORT=7000
APP_READ_TIMEOUT=30
APP_DEBUG=false

# JWT settings:
JWT_SECRET_KEY="super_secret_here"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=1440

# Database settings:
DB_HOST=niom-postgres
DB_PORT=5432
DB_USER=dev
DB_PASSWORD=dev
DB_NAME=niom_go_api
DB_SSL_MODE=disable
DB_DEBUG=true
DB_MAX_OPEN_CONNECTIONS=3
DB_MAX_IDLE_CONNECTIONS=1
DB_MAX_LIFETIME_CONNECTIONS=10
`