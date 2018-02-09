"use strict"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

const randomKey = () => {
  let key = ""
  for (let i = 0; i < 100; i++) {
    key += alphabet[Math.floor(Math.random() * 26)]
  }
  return key
}

const mod = (n, m) => (n % m + m) % m

let Cipher = function(userKey) {
  this.key = userKey || randomKey()
  if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
    throw Error("Bad key")
  }

  this.convert = (text, sign) => {
    let converted = ""
    Array.from(text).forEach((char, i) => {
      const shift = sign * alphabet.indexOf(this.key[mod(i, this.key.length)])
      converted +=
        alphabet[mod(alphabet.indexOf(char) + shift, alphabet.length)]
    })
    return converted
  }

  this.encode = (plaintext) => {
    return this.convert(plaintext, 1)
  }

  this.decode = (ciphertext) => {
    return this.convert(ciphertext, -1)
  }
}

module.exports = Cipher
