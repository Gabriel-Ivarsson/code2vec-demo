#!/usr/local/bin/python
import argparse
import gensim
from colors import green, red, blue
import sys

def check_word_list(package_name: str, word_list: list[str], input_model: str, model_file: str = "none"):
    if not word_list:
        # if list is empty
        print(red("Error: No words were provided as input"))
        sys.exit(1)

    if model_file == "none":
        if input_model == "fast":
            model = gensim.models.fasttext.FastText.load("./word2vec-model/fast-model.bin")
        elif input_model == "w2v":
            model = gensim.models.word2vec.Word2Vec.load("./word2vec-model/w2v-model.bin")
        else:
            print(red("No valid model was given as second argv argument"))
            sys.exit(1)
    else:
        if input_model == "fast":
            model = gensim.models.fasttext.FastText.load(model_file)
        elif input_model == "w2v":
            model = gensim.models.word2vec.Word2Vec.load(model_file)
        else:
            print(red("No valid model was given as second argv argument"))
            sys.exit(1)


    for word in word_list:
        print(blue(f"Results for words found similar to \"{word}\""))
        print(green(model.wv.most_similar(word)))

    for word in word_list:
        print(blue(f"Results for package name '{package_name}', compared to '{word}'"))
        print(green(model.wv.distance(package_name, word)))
    
    print(blue("Results for word that least fits in given word list"))
    print(red(model.wv.doesnt_match(word_list)))


def main():
    # command line arguments
    parser = argparse.ArgumentParser()
    parser.add_argument(dest = "function_list", nargs='+', default=[str], help="List of function names to compare to package name")
    parser.add_argument("-p", "--package",dest = "package_name", help="The Package Name to compare function list too", required=True)
    parser.add_argument("-m", "--model",dest = "input_model", help="The choosen input model; w2v or fast", required=True)
    parser.add_argument("-f", "--model-file",dest = "model_file", default="none", help="(Optional) Model file to load", required=False)
    args = parser.parse_args()

    # check against model
    check_word_list(args.package_name, args.function_list, args.input_model, args.model_file)



if __name__ == "__main__":
    main()
