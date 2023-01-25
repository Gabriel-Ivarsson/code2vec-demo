import gensim
from gensim.models.fasttext import FastText
from gensim.test.utils import datapath
import pandas as pd
import sys
import logging

if __name__ == '__main__':

    try:
        input_file = sys.argv[1]
        input_model = sys.argv[2]
    except Exception as e:
        print("Error: Supply model as second argv param, 'w2v' for Word2Vec or 'fast' FastText")
        sys.exit(1)

    if input_model == "fast":
        df = pd.read_json(input_file, lines=True)
        print(df)
        df.shape
        input_data = df.context.apply(gensim.utils.simple_preprocess)
        logging.info("Done reading data file!")

        # build vocabulary and train model
        model = FastText(
            input_data,
            vector_size=150,
            window=10,
            min_count=2,
            workers=10)

        model.build_vocab(input_data, progress_per=1000)
        model.train(input_data, total_examples=len(input_data), epochs=10)
        model.save("fast-model.bin")
    elif input_model == "w2v":
        df = pd.read_json(input_file, lines=True)
        print(df)
        df.shape
        input_data = df.context.apply(gensim.utils.simple_preprocess)
        logging.info("Done reading data file!")

        # build vocabulary and train model
        model = gensim.models.Word2Vec(
            window=10,
            min_count=1,
            workers=4,
            epochs=5,
        )

        model.build_vocab(input_data, progress_per=1000)
        model.train(input_data, total_examples=model.corpus_count, epochs=model.epochs)
        model.save("w2v-model.bin")
    else:
        print("Error: No valid model given")
        sys.exit(1)