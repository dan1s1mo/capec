def filter_by_list(allowed_words, parsed_words):
    return [parsed_word for parsed_word in parsed_words if parsed_word["text"].lower().replace("_", " ").strip() in allowed_words ]