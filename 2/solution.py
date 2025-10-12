def process(path):
    text = read_page(path)
    text = spl(text)
    res = search(text, "<a")
    return res


def search(text, key):
    n = 0
    for i in text:
        if len(i) == len(key):
            if i == key:
                n += 1
    return n

def read_page(name):
    with open(name) as f:
        return f.read()

def spl(text):
    return text.split()


