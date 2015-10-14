(ns rna-transcription)

(def dna-map {
              "G" "C"
              "C" "G"
              "T" "A"
              "A" "U"
              })

(defn to-rna [x] ((dna-map x)) x)


