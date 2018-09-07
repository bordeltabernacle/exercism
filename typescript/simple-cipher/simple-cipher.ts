class SimpleCipher {
  private alphabet: string = `abcdefghijklmnopqrstuvwxyz`
  key: string

  constructor(userKey?: string) {
    this.key = userKey || this.randomKey()
  }

  private randomKey() {
    return Array.from(Array(100)).reduce(
      (key) => key + this.alphabet[Math.floor(Math.random() * 26)],
      ""
    )
  }

  encode(plaintext: string): string {
    return plaintext
  }

  decode(ciphertext: string): string {
    return ciphertext
  }
}

export default SimpleCipher
