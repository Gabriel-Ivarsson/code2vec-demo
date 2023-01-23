import gensim
import pandas as pd
import sys

input = sys.argv[1]
df = pd.read_json(input, lines=True)
print(df)
df.shape

review_text = df.context.apply(gensim.utils.simple_preprocess)

model = gensim.models.Word2Vec(
    window=10,
    min_count=1,
    workers=4,
)

model.build_vocab(review_text, progress_per=1000)


model.train(review_text, total_examples=model.corpus_count, epochs=model.epochs)

model.save("model.bin")
