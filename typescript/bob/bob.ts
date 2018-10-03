class Bob {
  private isQuiet(dialogue: string): boolean {
    return dialogue === ``
  }

  private isYelling(dialogue: string): boolean {
    return (
      dialogue === dialogue.toUpperCase() &&
      dialogue.toUpperCase() !== dialogue.toLowerCase()
    )
  }

  private isQuestion(dialogue: string): boolean {
    return dialogue.endsWith(`?`)
  }

  hey(dialogue: string): string {
    const trimmed: string = dialogue.trim()
    if (this.isQuiet(trimmed)) {
      return `Fine. Be that way!`
    }
    if (this.isYelling(trimmed)) {
      return `Whoa, chill out!`
    }
    if (this.isQuestion(trimmed)) {
      return `Sure.`
    }
    return `Whatever.`
  }
}

export default Bob
