
def words_check(text):
    res = clean_text(text)
    res = capitalize_words(res)
    res = res_text(res)
    return res

def capitalize_words(text):
    text = text.split()
    res = ""
    for i in text:
        res += i[0].upper() + i[1:].lower() + " "

    return res

def clean_text(text):
    words = text.split()
    res = []
    for word in words:
        total = len(word)
        non_alpha = 0

        for ch in word:
            if not ch.isalpha():
                non_alpha += 1

        if non_alpha < total / 2:
            clean_word = "".join(ch for ch in word if ch.isalpha())
            if clean_word:
                res.append(clean_word)

    return " ".join(res)

def res_text(text):
    words = text.split()
    dic = {}
    for word in words:
        dic[word] = dic.get(word, 0) + 1
    return dic
