import sys

sys.setrecursionlimit(100000000)

def getSequenceLength(word,i):
    count = 1
    while((i+count) < len(word) and word[i+count] == word[i]):
        print("Failure in sequence")
        count += 1
    print("Finished sequence")
    return count

def getNewWord(word,i):
    print("started getWord")
    if i >= len(word):
        return ""
    count = getSequenceLength(word,i)
    print("Failure in new word at depth " + str(i))

    return str(count) + str(word[i]) + str(getNewWord(word,i+count))


def main():
    word = "3113322113"
    for i in range(0,40,1):
        word = getNewWord(word,0)
        print(word)
    print(len(word))

main()
