#!/usr/bin/env bash

set -euo pipefail

# Update helm dependencies and builds charts for sub-charts.

charts_path="../charts"
local_charts=(chainlink-cluster)

for chart in "${local_charts[@]}"; do
  echo "Building chart for $chart from $charts_path/$chart/Chart.lock"
  helm dependency build "$charts_path/$chart"
done
