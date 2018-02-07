class Pangram {
    private sentence: string
    private alphabet: string[] = Array.from("abcdefghijklmnopqrstuvwxyz")

    constructor(sentence: string) {
        this.sentence = sentence
    }

    isPangram(): boolean {
        if (this.sentence.length < 26) {
            return false
        }
        return this.alphabet.every((letter: string) =>
            this.sentence.toLowerCase().includes(letter)
        )
    }
}

export default Pangram
