class Pangram {
    private sentence: string[]
    private alphabet: string[] = Array.from("abcdefghijklmnopqrstuvwxyz")

    constructor(sentence: string) {
        this.sentence = Array.from(sentence)
    }

    isPangram(): boolean {
        if (this.sentence.length < 26) {
            return false
        }
        const sentenceMap = new Map<string, boolean>()
        this.sentence.forEach((key: string) => {
            sentenceMap.set(key.toLowerCase(), true)
        })
        return this.alphabet.every(
            (letter: string) => sentenceMap.get(letter) !== undefined
        )
    }
}

export default Pangram
