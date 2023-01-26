import gensim
from colors import cyan, green, red, blue
import sys

def check_word_list(input_list: list[str], input_model: str):
    if not input_list:
        # if list is empty
        print(red("Error: No words were provided as input"))
        sys.exit(1)

    if input_model == "fast":
        model = gensim.models.fasttext.FastText.load("./word2vec-model/fast-model.bin")
    elif input_model == "w2v":
        model = gensim.models.word2vec.Word2Vec.load("./word2vec-model/w2v-model.bin")
    else:
        print(red("No valid model was given as second argv argument"))
        sys.exit(1)

    for word in input_list:
        print(blue("Results for words found similar to \"" + word + "\""))
        print(green(model.wv.most_similar(word)))

    
    print(blue("Results for word that least fits in given word list"))
    print(red(model.wv.doesnt_match(input_list)))


def main(argv: list[str]):
    try:
        input_model = argv[1]
    except Exception as e:
        print("Error: upply model as second argv param, 'w2v' for Word2Vec or 'fast' FastText")
        raise e
    input_list = argv[2:] # get all arguments, minus file name and model

    # check against model
    check_word_list(input_list, input_model)



if __name__ == "__main__":
    main(sys.argv)
