class Bob {
  hey(dialogue: string): string {
    const s: string = dialogue.trim()
    if (s === "") {
      return "Fine. Be that way!"
    }
    if (s === s.toUpperCase() && s.toUpperCase() !== s.toLowerCase()) {
      return "Whoa, chill out!"
    }
    if (s.endsWith("?")) {
      return "Sure."
    }
    return "Whatever."
  }
}

export default Bob
