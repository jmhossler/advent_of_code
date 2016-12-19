def count(string1,string2):
    count = 0
    if(len(string1) != 2):
        return count
    for i in range(0,len(string2)-2,1):
        if(string1 == string2[i:i+2]):
            count += 1
            i += 2
    return count


def isNice(word):
    """strings = ['ab','cd','pq','xy']
    notBad = True
    vowels = ['a','e','i','o','u']
    vowelCount = 0
    prev = ''
    twice = False"""
    pair = False
    oneLet = False

    for i in range(0,len(word)-2,1):
        if(word[i-2] == word[i]):
            oneLet = True
        if(count(word[i:i+2],word) >= 2):
            pair = True

        """
        if(word[i] in vowels):
            vowelCount += 1
        if(word[i] == prev):
            twice = True
        if((prev + word[i]) in strings):
            notBad = False
        prev = word[i]
        """

    return pair and oneLet
    """
    return (notBad and (vowelCount >= 3) and twice)
    """


def main():
    f = open('input','r')
    word = f.readline()
    count = 0

    while(word != ''):
        if(isNice(word)):
            count += 1
        word = f.readline()

    print("Nice: {0}".format(count))

    return

main()
