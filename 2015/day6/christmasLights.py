def brightnessControl(lights,start,stop,inc):
    for i in range(int(start[0]),int(stop[0])+1,1):
        for j in range(int(start[1]),int(stop[1])+1,1):
            lights[i][j] += inc
            if lights[i][j] < 0:
                lights[i][j] = 0

def toggle(lights,start,stop):
    for i in range(int(start[0]),int(stop[0])+1,1):
        for j in range(int(start[1]),int(stop[1])+1,1):
            if lights[i][j] == 1:
                lights[i][j] = 0
            else:
                lights[i][j] = 1


def turn(lights,start,stop,onOff):
    for i in range(int(start[0]),int(stop[0])+1,1):
        for j in range(int(start[1]),int(stop[1])+1,1):
            lights[i][j] = onOff

def main():
    lights = []
    for i in range(0,1000):
        lights.append([])
        for j in range(0,1000):
            lights[i].append(0)
    fp = open("input","r")

    line = fp.readline()

    while(line != ''):
        line = line.split()
        if("toggle" in line):
            #print("Toggle: " +line[1] +" "+ line[3])
            #toggle(lights,line[1].split(','),line[3].split(','))
            brightnessControl(lights,line[1].split(','),line[3].split(','),2)
        elif("on" in line):
            #print("On: " +line[2] +" "+ line[4])
            #turn(lights,line[2].split(','),line[4].split(','),1)
            brightnessControl(lights,line[2].split(','),line[4].split(','),1)
        elif("off" in line):
            #print("Off: " +line[2] +" "+ line[4])
            #turn(lights,line[2].split(','),line[4].split(','),0)
            brightnessControl(lights,line[2].split(','),line[4].split(','),-1)
        line = fp.readline()

    fp.close()
    count = 0
    for array in lights:
        for value in array:
            count += value

    print("Brightness: " + str(count))

main()
