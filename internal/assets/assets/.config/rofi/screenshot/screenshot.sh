#!/bin/bash

options="Selection\x00icon\x1f/$HOME/.config/rofi/screenshot/selection.svg\nScreen\x00icon\x1f/$HOME/.config/rofi/screenshot/screen.svg"

choice=$(echo -e "$options" | rofi -dmenu -theme ~/.config/rofi/screenshot/screenshot.rasi)

if [ -n "$choice" ]; then
    case "$choice" in
        Selection)
            (screenshot -g "$(slurp -w 0)" - | wl-copy) &
            KEEP_PID=$!
            ;;
        Screen)
            (sleep 0.5; screenshot - | wl-copy) &
            KEEP_PID=$!
            ;;
        *)
            exit 1
            ;;
    esac
fi