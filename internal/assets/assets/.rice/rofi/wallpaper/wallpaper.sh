#!/bin/bash

DIR="$HOME/.rice/wallpapers"
prompt="Wallpaper"

rofi_cmd() {
	rofi -dmenu \
		-p "$prompt" \
		-markup-rows \
		-theme "~/.rice/rofi/wallpaper/wallpaper.rasi"
}

run_rofi() {
    while IFS= read -r -d '' file; do
        fname="$(basename "$file")"
        printf "%s\0icon\x1f%s\n" "$fname" "$file"
    done < <(find "$DIR" -type f -print0) | rofi_cmd
}

wallpaper="$(run_rofi)"

echo $wallpaper

if [ -n "$wallpaper" ]; then
    swww img "$DIR/$wallpaper" -t wipe --transition-duration 2 --transition-fps 60
fi