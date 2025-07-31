# Variables
BINARY_NAME         := "hgmx"
MAIN_PACKAGE_PATH   := "./cmd/hgmx"
VERSION_VAR_PATH    := "main.Version"
GIT_VERSION         := `git describe --tags --always || echo dev`
CSS_INPUT           := "views/static/css/main.css"
CSS_OUTPUT          := "views/static/css/main.min.css"
PATH_SCRIPTS        := "views/static/scripts"

default:
    @just --list

gen:
    @templ generate

templ-update:
    go install github.com/a-h/templ/cmd/templ@latest

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
    @tailwindcss -i {{CSS_INPUT}} -o {{CSS_OUTPUT}} --minify

watch:
    @air

describe:
    @echo "{{GIT_VERSION}}"

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