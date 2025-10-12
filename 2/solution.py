import re

def process(path):
    text = read_page(path)
    text = clean_html(text)
    res = search(text)
    return res


def clean_html(text):
    text = re.sub(r'<!--.*?-->', '', text, flags=re.DOTALL)
    text = re.sub(r'<script.*?>.*?</script>', '', text, flags=re.DOTALL | re.IGNORECASE)
    text = re.sub(r'<style.*?>.*?</style>', '', text, flags=re.DOTALL | re.IGNORECASE)
    return text


def search(text):
    matches = re.findall(r'<\s*a\b[^>]*>', text, flags=re.IGNORECASE)
    return len(matches)


def read_page(name):
    with open(name, encoding="utf-8") as f:
        return f.read()


def spl(text):
    return text.split()
