(ns bob
  (:require [clojure.string :as str]))

(defn response-for
  [spoken]
  (cond
    (and (= (.toUpperCase spoken) spoken)
         (some #(Character/isLetter %) spoken)) "Whoa, chill out!"
    (str/blank? spoken)                         "Fine. Be that way!"
    (= (last spoken) \?)                        "Sure."
    :else                                       "Whatever."))

