import gensim
import sys
input = sys.argv[1]
model = gensim.models.Word2Vec.load("model.bin")

print("Results for words found similar to \"" + input + "\"")
print(model.wv.most_similar(input))
