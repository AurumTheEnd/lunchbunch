#!/bin/bash

set -e

cd "$(dirname "$0")/.."

mkdir -p /tmp/gitlab.fi.muni.cz/xhrdlic3/lunchbunch
cp -R . /tmp/gitlab.fi.muni.cz/xhrdlic3/lunchbunch

for resolution in "file" "package"; do
  goimportcycle -dot ./diagrams/package.dot -resolution $resolution -path /tmp/gitlab.fi.muni.cz/xhrdlic3/lunchbunch
  dot -Tpng -o ./diagrams/package-$resolution.png ./diagrams/package.dot
done

rm -rf /tmp/gitlab.fi.muni.cz/xhrdlic3/lunchbunch