#!/bin/bash
source "$(dirname "$0")/.env"

smee -u "$WEBHOOK_PROXY_URL" -p 8080