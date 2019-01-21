#!/usr/bin/env bash

set -euo pipefail

obxmodels=internal/pkg/db/objectbox/models #source
obxbindings=internal/pkg/db/objectbox/obx  #target

generator="go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen/ -byValue -persist ${obxbindings}/objectbox-model.json"

# generate
for f in ${obxmodels}/*.go; do ${generator} -source ${f}; done

# fix import path
for f in ${obxmodels}/*obx.go; do sed -i 's/import (/import (\n\t. "github.com\/edgexfoundry\/edgex-go\/pkg\/models"/g' "$f"; done

# move to the output folder
mv ${obxmodels}/*obx.go ${obxbindings}/

# fix package name on generated files and objectbox-model.go
for f in ${obxbindings}/*.go; do sed -i 's/package models/package obx/g' "$f"; done