(ns bob
  (:require [clojure.string :as str]))

(defn has-letter? [sentence] (some #(Character/isLetter %) sentence))
(defn yelling?    [sentence] (and (= (str/upper-case sentence) sentence)
                                  (has-letter? sentence)))
(defn question?   [sentence] (and (= \? (last sentence))
                                  (not (yelling? sentence))))
(defn silence?    [sentence] (str/blank? sentence))

(defn response-for
  [conversation]
  (cond
    (silence?   conversation) "Fine. Be that way!"
    (question?  conversation) "Sure."
    (yelling?   conversation) "Whoa, chill out!"
    :else "Whatever."))
