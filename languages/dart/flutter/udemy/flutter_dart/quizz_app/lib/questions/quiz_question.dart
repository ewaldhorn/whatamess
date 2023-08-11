class QuizQuestion {
  final String question;
  final List<String> answers;

  const QuizQuestion(this.question, this.answers);

  List<String> getShuffledAnswers() {
    final newList = List.of(answers);
    newList.shuffle();
    return newList;
  }
}
