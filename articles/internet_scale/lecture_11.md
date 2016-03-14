Lecture 11
===========

### Classical view of Search
Documents have keywords. We query the document for keywords

Search process is two-part: indexing and querying

Indexing Algo:

    index = dict()  # keyword -> document
    for d in documents:
      for kw in d:
        index[kw] = d

Querying Algo:

    results = set()
    for kw in queries:
      results += index[kw]
    return results


### Search in 1970s
Not all keywords are equal.
  * do not want to query for 'the'

Thus, add ranking to the search process. This

Term frequency-i...(TF-IDF)
  * words popularity is important in determining score
  * rare words have higher score


Indexing Algo:

    # score() - gives a score for a word
    index = dict()
    for d in documents:
      for kw in d:
        index[kw] = (d,score(kw,d))

Querying Algo:

    results = dict()
    for kw in queries:
      d,score=index[kw]
      results[d] += score
    return sorted(results.items(), key=lambda x: x[1])

1970s score function:

    # map from (kw,d) -> count of times kw appears in d
    tf = dict()
    # map from kw to what fraction of all known documents contain kw
    df = dict()
    def score(kw, d):
      # consider both whether word rare or not in either a given doc or in general
      return tf(kw,d) * (1.0-df(kw))


For example:

    d1="the book was good"
    d2="the book gilgamesh is old"
    query="gilgamesh book"

    score d1 for this query:
      score("gilgamesh", d1)+score("book", d1)
      =tf["gilgamesh", d1] * (1-df("gilgamesh")) + tf["book", d1] * (1-df["book"])
      =0 * (1-0.5) + 1 * (1-1)
      =0

    score d2 for this query:
      score("gilgamesh", d2)+score("book", d2)
      =tf["gilgamesh", d2] * (1-df("gilgamesh")) + tf["book", d2] * (1-df["book"])
      =1 * (1-0.5) + 1 * (1-1)
      =0.5
