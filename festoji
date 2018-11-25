#!/bin/bash
set -euo pipefail

DAY=$((60 * 60 * 24))
TODAY=$(date --date="today 00:00:00" +%s)

in_season() {
    # Prints the given emoji if TODAY is within date range.
    # End date range is non-inclusive.
    local TODAY=$1
    local end=$2
    local range=$3
    local emoji=$4

    local season=$(($DAY * $range))
    local start=$(($end - $season))

    if [ "$TODAY" -ge "$start" -a "$TODAY" -lt "$end" ]; then
        echo $emoji
        exit 0
    fi
}

nth_weekday() {
    local month=$1
    local expected_weekday=$2  # Sunday is 0
    # Third parameter is nth week in month, first week is 0
    # The expression below computes the day of the month that always
    # appears on the nth week, e.g. 1st on week 0, 8th on week 1
    local n=$((($3 * 7) + 1))

    local beg beg_weekday
    read beg beg_weekday <<< $(date --date="${month} ${n} 00:00:00" '+%s %w')
    local final=$((
        $beg + ( ($expected_weekday - $beg_weekday) * DAY)
    ))
    printf $final
}

xmas_end=$(date --date="december 26 00:00:00" +%s)
in_season $TODAY $xmas_end 14 🎄

# The number 5 means Friday, which is when Thanksgiving ends.
# The number 3 means the fourth week of the month.
thanksgiving_end=$(nth_weekday 'november' 5 3)
in_season $TODAY $thanksgiving_end 7 🦃

newyears_end=$(date --date="january 1 00:00:00" +%s)
in_season $TODAY $newyears_end 5 🍾

echo 🐚
