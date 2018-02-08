class Pangram {
    private text: string[]
    constructor(text: string) {
        this.text = Array.from(text)
    }

    private textMap(): Map<string, boolean> {
        const m = new Map<string, boolean>()
        this.text.forEach((key: string) => {
            m.set(key.toLowerCase(), true)
        })
        return m
    }

    isPangram(): boolean {
        if (this.text.length < 26) {
            return false
        }
        for (let letter = 97; letter < 123; letter++) {
            if (this.textMap().get(String.fromCharCode(letter)) === undefined) {
                return false
            }
        }
        return true
    }
}

export default Pangram
