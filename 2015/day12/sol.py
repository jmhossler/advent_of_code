# Advent of Code: Day 12
import sys
import re
import json

f = open(sys.argv[1],"r")
string = f.read()
f.close()

# Part 1
print "Sum of all numbers 1:", sum(map(int, re.findall("-?[0-9]+", string)))

# Part 2
def hook(obj):
  if "red" in obj.values(): return {}
  else: return obj
stuff = str(json.loads(string, object_hook=hook))
print "Sum of all numbers 2:", sum(map(int, re.findall("-?[0-9]+", stuff)))
