in=3113322113;
for f in $(seq 1 50); do
  in=$(echo "$in" | fold -w1 | uniq -c | tr '\n' ' ' | tr -d ' ');
  echo $in | tr -d '\n' | wc -c;
done
