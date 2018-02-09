class Cipher {
  constructor(userKey) {
    this.alphabet = "abcdefghijklmnopqrstuvwxyz"
    this.key = userKey || this.randomKey()
    if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
      throw Error("Bad key")
    }
  }

  randomKey() {
    let key = ""
    for (let i = 0; i < 100; i++) {
      key += this.alphabet[Math.floor(Math.random() * 26)]
    }
    return key
  }

  _convertedLetter(letter, index, sign) {
    let offset =
      this.alphabet.indexOf(letter) +
      sign * this.alphabet.indexOf(this.key[index % this.key.length])
    offset = offset >= 26 ? (offset -= 26) : offset
    offset = offset < 0 ? (offset += 26) : offset
    return this.alphabet[offset]
  }

  _convert(text, sign) {
    let converted = ""
    Array.from(text).forEach((letter, index) => {
      converted += this._convertedLetter(letter, index, sign)
    })
    return converted
  }

  encode(plaintext) {
    return this._convert(plaintext, 1)
  }

  decode(ciphertext) {
    return this._convert(ciphertext, -1)
  }
}

module.exports = Cipher
