class Cipher {
  constructor(userKey) {
    this.alphabet = "abcdefghijklmnopqrstuvwxyz"
    this.key = userKey || this.randomKey()
    if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
      throw Error("Bad key")
    }
  }

  randomKey() {
    return Array.from(Array(100)).reduce(
      (key) => key + this.alphabet[Math.floor(Math.random() * 26)],
      ""
    )
  }

  _wrap(offset) {
    offset = offset >= 26 ? (offset -= 26) : offset
    offset = offset < 0 ? (offset += 26) : offset
    return offset
  }

  _convertedLetter(letter, index, sign) {
    let offset =
      this.alphabet.indexOf(letter) +
      sign * this.alphabet.indexOf(this.key[index % this.key.length])
    return this.alphabet[this._wrap(offset)]
  }

  _convert(text, sign) {
    return Array.from(text).reduce((converted, letter, index) => {
      return converted + this._convertedLetter(letter, index, sign)
    }, "")
  }

  encode(plaintext) {
    return this._convert(plaintext, 1)
  }

  decode(ciphertext) {
    return this._convert(ciphertext, -1)
  }
}

module.exports = Cipher
