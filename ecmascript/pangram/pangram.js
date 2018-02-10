class Pangram {
  constructor(text) {
    this.text = text
  }

  isPangram() {
    return (
      new Set([...this.text.toLowerCase().replace(/[^a-z]/g, "")]).size === 26
    )
  }
}

module.exports = Pangram
