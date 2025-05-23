#!/usr/bin/env bash

# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# cd to the root path
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${ROOT}"
PATH="${PATH}:${ROOT}/bin:$(go env GOPATH)/bin"

echo "Updating tables of contents if necessary..."
# Update tables of contents if necessary.
find proposals -mindepth 2 -name '*.md' \
    | xargs mdtoc --inplace --max-depth=5  || (
      echo "Failed generating TOC. If this failed silently and you are on mac, try 'brew install grep'"
      exit 1
    )
