class SimpleCipher {
  private alphabet: string = `abcdefghijklmnopqrstuvwxyz`
  private alphabetLength: number = this.alphabet.length

  readonly key: string

  constructor(userKey?: string) {
    this.key = userKey || this.randomKey()
    if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
      throw Error("Bad key")
    }
  }

  private randomKey(): string {
    return Array.from(Array(100)).reduce(
      (key: string) =>
        `${key}${
          this.alphabet[Math.floor(Math.random() * this.alphabetLength)]
        }`,
      ""
    )
  }

  private adjustOffset(offset: number): number {
    if (offset >= this.alphabetLength) {
      offset -= this.alphabetLength
    }
    if (offset < 0) {
      offset += this.alphabetLength
    }
    return offset
  }

  private convertChar(char: string, index: number, sign: number): string {
    const offset: number = this.adjustOffset(
      this.alphabet.indexOf(char) +
        sign * this.alphabet.indexOf(this.key[index % this.key.length])
    )
    return this.alphabet[offset]
  }

  private convert(text: string, sign: number) {
    return Array.from(text).reduce((result, char, i) => {
      return `${result}${this.convertChar(char, i, sign)}`
    }, "")
  }

  encode(plaintext: string): string {
    return this.convert(plaintext, 1)
  }

  decode(ciphertext: string): string {
    return this.convert(ciphertext, -1)
  }
}

export default SimpleCipher
