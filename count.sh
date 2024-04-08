total=$(find -mindepth 0 | wc -l)
total=$(( total * 5 ))
printf "\t\v%s\v\n" "Total files * 5: $total"