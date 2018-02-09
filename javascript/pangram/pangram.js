let Pangram = function(text) {
  this.text = text

  this.isPangram = () => {
    return (
      new Set(Array.from(this.text.toLowerCase().replace(/[^a-z]/g, "")))
        .size === 26
    )
  }
}

module.exports = Pangram
