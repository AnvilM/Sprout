#!/bin/bash

battery=$(ls /sys/class/power_supply/ | grep -i '^BAT' | head -n1)

if [ -z "$battery" ]; then
    ~/.rice/notifications/error/error.sh "No battery found in /sys/class/power_supply!"
    exit 1
fi

battery_level=$(< /sys/class/power_supply/"$battery"/capacity)
battery_status=$(< /sys/class/power_supply/"$battery"/status)

if [ "$battery_status" = "Discharging" ]; then
    if [ "$battery_level" -le 10 ]; then
        notify-send -u critical -i ~/.rice/notifications/battery/icon_battery_critical.svg "Battery Critical" "${battery_level}% left"
    elif [ "$battery_level" -le 20 ]; then
        notify-send -i ~/.rice/notifications/battery/icon_battery_low.svg "Battery Low" "${battery_level}% left"
    fi
fi
