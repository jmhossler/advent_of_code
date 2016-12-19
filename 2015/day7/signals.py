def notValue(dic,line):
    if(line[1].isdigit()):
        dic[line[len(line)-1]] = ~(int(line[1]))
    else:
        dic[line[len(line)-1]] = line[0:2]

def process(dic,line):
    dic[line[len(line)-1]] = line[0:len(line)-2]

def findAnswer(dic,key):
    line = dic.get(key)
    if(type(line) == 'int'):
        return int(line)
    x = 0
    y = 0
    if 'NOT' in line:
        x = int(findAnswer(dic,line[1]))
    else:
        x = findAnswer(dic,line[0])
        y = findAnswer(dic,line[2])
    if 'NOT' in line:
        dic[key] = ~x
    if 'RSHIFT' in line:
        dic[key] = x >> y
    if 'LSHIFT' in line:
        dic[key] = x << y
    if 'AND' in line:
        dic[key] = x & y
    if 'OR' in line:
        dic[key] = x | y

    return dic[key]

def main():
    fp = open("input","r")

    line = fp.readline()
    dic = {}

    while(line != ''):
        line = line.replace('\n','').split(' ')
        if('NOT' in line):
            notValue(dic,line)
        else:
            process(dic,line)

        line = fp.readline()
    for key in dic.keys():
        x = findAnswer(dic,key)

    print(dic)

main()
