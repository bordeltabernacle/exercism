class SimpleCipher {
  private alphabet: string = `abcdefghijklmnopqrstuvwxyz`
  private key: string

  constructor(userKey?: string) {
    this.key = userKey || this.randomKey()
    if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
      throw Error("Bad key")
    }
  }

  private randomKey(): string {
    return Array.from(Array(100)).reduce(
      (key: string) => `${key}${this.alphabet[Math.floor(Math.random() * 26)]}`,
      ""
    )
  }

  private convertChar(char: string, index: number, sign: number): string {
    let offset: number =
      this.alphabet.indexOf(char) +
      sign * this.alphabet.indexOf(this.key[index % this.key.length])

    offset = offset >= 26 ? (offset -= 26) : offset
    offset = offset < 0 ? (offset += 26) : offset

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
