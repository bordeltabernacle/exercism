class Bob {
  hey(dialogue: string): string {
    const trimmed = dialogue.trim();
    if (trimmed === ``) {
      return `Fine. Be that way!`;
    }
    if (
      trimmed === trimmed.toUpperCase() &&
      trimmed.toUpperCase() !== trimmed.toLowerCase()
    ) {
      return `Whoa, chill out!`;
    }
    if (trimmed.endsWith(`?`)) {
      return `Sure.`;
    }
    return `Whatever.`;
  }
}
export default Bob;
