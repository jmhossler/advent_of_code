COL = 3019
ROW = 3010
START = 20151125
MULT = 252533
MOD = 33554393

def getVal(num):
    return (num * MULT) % MOD

def fillDiag(array,start):
    array[start][0] = getVal(array[0][start-1])

    val = array[start][0]
    i = 1
    start = start - 1

    while start >= 0:
        array[start][i] = getVal(val)
        val = array[start][i]
        start = start - 1
        i = i + 1

def display(array):
    for i in range(0,7):
        for j in range(0,7):
            print(array[i][j],end=' ')
        print()


def main():
    array = [[0 for i in range(10000)] for j in range(10000)]

    array[0][0] = START

    i = 1

    while(array[ROW+1][COL+1] == 0):
        display(array)
        fillDiag(array,i)
        i = i + 1

    display(array)

    print(array[ROW+1][COL+1])

main()
