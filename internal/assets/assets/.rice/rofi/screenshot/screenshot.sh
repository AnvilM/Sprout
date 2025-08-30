#!/bin/bash


prompt='Screenshot'

option_1="󰩭 Capture Area"
option_2="󱣴 Capture Screen"
option_3=" Capture in 5s"
option_4=" Capture in 10s"

rofi_cmd() {
	rofi -dmenu \
		-p "$prompt" \
		-markup-rows \
		-theme "~/.rice/rofi/screenshot/screenshot.rasi"
}

run_rofi() {
	echo -e "$option_1\n$option_2\n$option_3\n$option_4" | rofi_cmd
}

shotscreen() {
    sleep 0.5
    ~/.rice/scripts/screenshot.sh
}

shotarea() {
    ~/.rice/scripts/screenshot.sh -g "$(slurp -w 0)"
}

shot5() {
    sleep 5
    ~/.rice/scripts/screenshot.sh
}

shot10() {
    sleep 10
    ~/.rice/scripts/screenshot.sh
}

run_cmd() {
	if [[ "$1" == '--screen' ]]; then
		shotscreen
	elif [[ "$1" == '--area' ]]; then
		shotarea
	elif [[ "$1" == '--shot5' ]]; then
		shot5
	elif [[ "$1" == '--shot10' ]]; then
		shot10
	fi
}

chosen="$(run_rofi)"
case ${chosen} in
    $option_1)
		run_cmd --area
        ;;
    $option_2)
		run_cmd --screen
        ;;
    $option_3)
		run_cmd --shot5
        ;;
    $option_4)
		run_cmd --shot10
        ;;
    *)
        exit 1
        ;;
esac