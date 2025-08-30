#!/bin/bash
set -o pipefail

take_screenshot() {
    grim "$@" - | wl-copy
}

if take_screenshot "$@"; then
    ~/.rice/notifications/screenshot/screenshot.sh
fi