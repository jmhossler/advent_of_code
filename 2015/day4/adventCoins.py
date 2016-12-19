import hashlib

def main():
    string = "iwrupvqb".encode('utf-8')
    m = hashlib.md5()
    m.update(string)
    h = m.copy()
    for i in range(0,1000000000,1):
        h.update(str(i).encode('utf-8'))
        string = h.hexdigest()
        if string[0:6] == "000000":
            print(i)
            break
        h = m.copy()
    return


main()
