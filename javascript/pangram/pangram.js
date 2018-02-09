let Pangram = function(text) {
  this.text = text

  this.isPangram = () => {
    if (this.text.length < 26) {
      return false
    }
    for (let letter = 97; letter < 123; letter++) {
      if (!this.text.toLowerCase().includes(String.fromCharCode(letter))) {
        return false
      }
    }
    return true
  }
}

module.exports = Pangram
