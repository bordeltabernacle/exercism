def word_count(string):
    word_dict = {}
    word_list = string.split()
    for word in word_list:
        if word not in word_dict:
            word_dict[word] = 1
        else:
            word_dict[word] += 1
    return word_dict
