"use strict"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

const randomKey = () => {
  let key = ""
  for (let i = 0; i < 100; i++) {
    key += alphabet[Math.floor(Math.random() * 26)]
  }
  return key
}

module.exports = function(userKey) {
  this.key = userKey || randomKey()
  if (userKey === "" || !this.key.match(/^[a-z]+$/)) {
    throw Error("Bad key")
  }

  this.encode = plaintext => {
    let encoded = ""
    plaintext.split("").forEach((char, i) => {
      let shift =
        alphabet.indexOf(char) + alphabet.indexOf(this.key[i % this.key.length])
      if (shift >= 26) {
        shift -= 26
      }
      encoded += alphabet[shift]
    })
    return encoded
  }

  this.decode = ciphertext => {
    let decoded = ""
    ciphertext.split("").forEach((char, i) => {
      let shift =
        alphabet.indexOf(char) - alphabet.indexOf(this.key[i % this.key.length])
      if (shift < 0) {
        shift += 26
      }
      decoded += alphabet[shift]
    })
    return decoded
  }
}
