def calculate_word_score(word, order_dict):
    score = 0
    for char in word:
        if char in order_dict:
            score += order_dict[char]
    return score


def queranumeric(order: list[str], words: list[str]) -> list[str]:
    order_dict = {}
    for i in range(len(order)):
        order_dict[order[i]] = i

    def compare_key(word):
        result = []
        for char in word:
            if char in order_dict:
                result.append(order_dict[char])
            else:
                result.append(len(order))
        return result

    return sorted(words, key=compare_key)


def main():
    result = queranumeric(list("namrepus"), ["sun", "man", "super", "name", "user", "spam", "ram"])
    print("نتیجه:", result)
    print("انتظار: ['name', 'man', 'ram', 'user', 'spam', 'sun', 'super']")
    print("درست؟", result == ['name', 'man', 'ram', 'user', 'spam', 'sun', 'super'])

main()