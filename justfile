# Variables
BINARY_NAME         := "hgmx"
MAIN_PACKAGE_PATH   := "./cmd/hgmx"
VERSION_VAR_PATH    := "main.Version"
GIT_VERSION         := `git describe --tags --always || echo dev`
CSS_INPUT           := "views/static/css/main.css"
CSS_OUTPUT          := "views/static/css/main.min.css"
PATH_SCRIPTS        := "views/static/scripts"
COMPOSE             := "docker compose"
DB_SERVICE          := "hgmx-pg"
WATCHER_PORT        := "8003"

default:
    @just --list

gen:
    @templ generate

templ-update:
    go install github.com/a-h/templ/cmd/templ@latest

air-update:
    go install github.com/air-verse/air@latest

upgrade:
    @just templ-update
    @just air-update
    go get -u ./...
    go mod tidy

build: gen
    @go build -ldflags="-X {{VERSION_VAR_PATH}}={{GIT_VERSION}}" -o {{BINARY_NAME}} {{MAIN_PACKAGE_PATH}}

install:
    @go install ./cmd/{{BINARY_NAME}}

clean:
    @rm -f {{BINARY_NAME}}

# e.g., run info -v
run *ARGS:
    @go run {{MAIN_PACKAGE_PATH}} {{ARGS}}

tw:
    @tailwindcss -i {{CSS_INPUT}} -o {{CSS_OUTPUT}} --minify --watch --verbose

watch:
    #!/usr/bin/env bash
    set -euo pipefail

    if ! docker info >/dev/null 2>&1; then
        echo "Docker is not running. Please start Docker first."
        exit 1
    fi

    if ! {{COMPOSE}} ps --services --filter "status=running" | grep -q "{{DB_SERVICE}}"; then
        {{COMPOSE}} up -d
        if {{COMPOSE}} help 2>/dev/null | grep -q " wait "; then
            {{COMPOSE}} wait || true
        else
            sleep 2
        fi
    fi

    env PROXY=true \
    templ generate --watch --proxy="http://localhost:3008" --proxyport={{WATCHER_PORT}} --open-browser=false --cmd="go run cmd/main.go" --log-level=warn &
    TEMPL_PID=$!

    cleanup() {
        kill $TEMPL_PID 2>/dev/null || true
        wait $TEMPL_PID 2>/dev/null || true
        exit 0
    }

    trap cleanup SIGINT SIGTERM

    air
    cleanup

patch:
    #!/usr/bin/env bash
    set -euo pipefail
    # Check for uncommitted changes
    if [[ -n $(git status --porcelain) ]]; then
      echo "Error: Working directory is not clean. Please commit or stash changes."
      exit 1
    fi
    
    LATEST_TAG=$(git describe --tags --abbrev=0)
    MAJOR_MINOR=$(echo $LATEST_TAG | awk -F. '{print $1"."$2}')
    CURRENT_PATCH=$(echo $LATEST_TAG | awk -F. '{print $3}')
    NEW_PATCH=$((CURRENT_PATCH + 1))
    NEW_VERSION_NUM="${MAJOR_MINOR#v}.${NEW_PATCH}"
    NEW_TAG="v${NEW_VERSION_NUM}"
    echo -n "$NEW_TAG" > ".version"
    git add ".version"
    git commit --amend --no-edit
    git tag $NEW_TAG

patch-undo:
    #!/usr/bin/env bash
    set -euo pipefail
    # Get the most recently created tag
    LATEST_TAG=$(git tag --sort=creatordate | tail -n 1)
    if [[ -z "$LATEST_TAG" ]]; then
      echo "Error: No tags found to delete."
      exit 1
    fi
    git tag -d $LATEST_TAG
    git push origin --delete $LATEST_TAG

release:
    @just patch 
    @git push
    @git push origin --tags

ds *URL:
    #!/usr/bin/env bash
    name=$(basename {{URL}})
    curl -o {{PATH_SCRIPTS}}/${name} {{URL}}
    if [ $? -eq 0 ]; then
        echo "downlaoded {{PATH_SCRIPTS}}/${name}"
    else
        echo "failed to download {{PATH_SCRIPTS}}/${name}"
        exit 1
    fi


new NAME:
    #!/usr/bin/env bash
    set -euo pipefail

    COMPONENT_PATH={{NAME}}
    PACKAGE_NAME=$(basename "${COMPONENT_PATH}")
    COMPONENT_DIR="views/components/${COMPONENT_PATH}"

    if [ -d "${COMPONENT_DIR}" ]; then
        echo "Error: Component directory already exists: ${COMPONENT_DIR}"
        exit 1
    fi

    mkdir -p "${COMPONENT_DIR}"

    {
        echo "package ${PACKAGE_NAME}"
        echo ""
        echo "func Method() *Props {"
        echo "	return &Props{}"
        echo "}"
    } > "${COMPONENT_DIR}/factories.go"

    {
        echo "package ${PACKAGE_NAME}"
        echo ""
        echo "type Props struct {"
        echo "	// components properties"
        echo "}"
    } > "${COMPONENT_DIR}/properties.go"

    {
        echo "package ${PACKAGE_NAME}"
        echo ""
        echo "templ (p *Props) Render() {"
        echo "	// html to render via templ"
        echo "}"
    } > "${COMPONENT_DIR}/render.templ"

    echo "Created component skeleton in ${COMPONENT_DIR}"

page NAME:
    #!/usr/bin/env bash
    set -euo pipefail

    PAGE_PATH={{NAME}}
    PAGE_DIR="views/pages/${PAGE_PATH}"

    if [ -d "${PAGE_DIR}" ]; then
        echo "Error: Page directory already exists: ${PAGE_DIR}"
        exit 1
    fi

    mkdir -p "${PAGE_DIR}"

    {
        echo "package ${PAGE_PATH}"
        echo ""
        echo "templ Main() {"
        echo "	<div>"
        echo "		<h1>${PAGE_PATH}</h1>"
        echo "	</div>"
        echo "}"
    } > "${PAGE_DIR}/${PAGE_PATH}.templ"

    echo "Created page skeleton in ${PAGE_DIR}"


db-up:
    @{{COMPOSE}} up -d {{DB_SERVICE}}

db-down:
    @{{COMPOSE}} down

db-restart:
    @{{COMPOSE}} restart {{DB_SERVICE}}

db-reset:
    @{{COMPOSE}} down
    @{{COMPOSE}} up -d {{DB_SERVICE}}

db-logs:
    @{{COMPOSE}} logs -f {{DB_SERVICE}}

db-url:
    #!/usr/bin/env sh
    set -eu
    SERVICE={{DB_SERVICE}}
    CID=$({{COMPOSE}} ps -q "$SERVICE" || true)
    if [ -z "${CID}" ]; then
        echo "${SERVICE} not running; starting..." >&2
        {{COMPOSE}} up -d "$SERVICE" >/dev/null
        CID=$({{COMPOSE}} ps -q "$SERVICE")
    fi
    host=127.0.0.1
    port=$({{COMPOSE}} port "$SERVICE" 5432 2>/dev/null | head -n1 | awk -F: '{print $NF}')
    if [ -z "${port:-}" ]; then port=5432; fi
    env_out=$({{COMPOSE}} exec -T "$SERVICE" env)
    user=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_USER=' | tail -n1 | cut -d= -f2)
    pass=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_PASSWORD=' | tail -n1 | cut -d= -f2)
    db=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_DB=' | tail -n1 | cut -d= -f2)
    user=${user:-postgres}
    pass=${pass:-postgres}
    db=${db:-postgres}
    echo "postgres://${user}:${pass}@${host}:${port}/${db}?sslmode=disable"

db-psql:
    #!/usr/bin/env sh
    set -eu
    SERVICE={{DB_SERVICE}}
    CID=$({{COMPOSE}} ps -q "$SERVICE" || true)
    if [ -z "${CID}" ]; then
        {{COMPOSE}} up -d "$SERVICE" >/dev/null
        CID=$({{COMPOSE}} ps -q "$SERVICE")
    fi
    env_out=$({{COMPOSE}} exec -T "$SERVICE" env)
    user=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_USER=' | tail -n1 | cut -d= -f2)
    pass=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_PASSWORD=' | tail -n1 | cut -d= -f2)
    db=$(printf "%s\n" "$env_out" | grep -E '^POSTGRES_DB=' | tail -n1 | cut -d= -f2)
    {{COMPOSE}} exec -e PGPASSWORD="${pass:-postgres}" "$SERVICE" psql -U "${user:-postgres}" -d "${db:-postgres}"