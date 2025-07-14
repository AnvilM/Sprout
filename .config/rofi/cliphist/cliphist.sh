#!/bin/bash

if [[ "$ROFI_RETV" -eq 0 ]]; then
    cliphist list | cut -f2-
elif [[ "$ROFI_RETV" -eq 1 ]]; then
    echo -n "$@​" | wl-copy
fi