import gensim
import pandas as pd
import sys
import logging

if __name__ == '__main__':

    try:
        input_file = sys.argv[1]
    except Exception as e:
        print("Supply file as argv param, or 'lazy' to use pre-downloaded data.")
        sys.exit(1)

    if input_file == "lazy":
        df = pd.read_json("./data/reviews_Cell_Phones_and_Accessories.json", lines=True)
        print(df)
        df.shape
        input_data = df.reviewText.apply(gensim.utils.simple_preprocess)
        logging.info("Done reading data file!")

        # build vocabulary and train model
        model = gensim.models.Word2Vec(
            input_data,
            vector_size=150,
            window=10,
            min_count=2,
            workers=10)

        model.build_vocab(input_data, progress_per=1000)
        model.train(input_data, total_examples=len(input_data), epochs=10)
    else:
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

    model.save("model.bin")
