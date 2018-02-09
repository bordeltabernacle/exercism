class Pangram {
  constructor(text) {
    this.text = text
    this._set = new Set(
      Array.from(this.text.toLowerCase().replace(/[^a-z]/g, ""))
    )
  }

  isPangram() {
    return this._set.size === 26
  }
}

module.exports = Pangram
