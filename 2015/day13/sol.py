f = open("input","r")
lines = f.read().split("\n")[:-1]
f.close()

data = dict()
for line in lines:
    parts = line.split(" ")
    name = parts[0]
    speed = int(parts[3])
    time = int(parts[6])
    rest = int(parts[-2])
    data[name] = (speed,time,rest)

names = data.keys()
distance = [0 for _ in names]
restingTimers = [0 for _ in names]
flyingTimers = [0 for _ in names]
states = [1 for _ in names]
points = [0 for _ in names]

timeCut = 2503
for t in range(timeCut):
    for i,name in enumerate(names):
        if states[i] == 1:
            distance[i] += data[name][0]
            flyingTimers[i] +=1
            if flyingTimers[i] == data[name][1]:
                flyingTimers[i] = 0
                states[i] = 0
        else:
            restingTimers[i] +=1
            if restingTimers[i] == data[name][2]:
                restingTimers[i] = 0
                states[i] = 1

    maxDistance = max(distance)
    for i,value in enumerate(distance):
        if value==maxDistance:
            points[i]+=1

print max(distance)
print max(points)
