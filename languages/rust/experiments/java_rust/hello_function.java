class HelloFunction {
    public static void main(String[] args) {
        HelloFunction.greetings();
    }

    public static void greetings() {
        boolean isCoding = true;
        String mood = (isCoding) ? "happy" : "sad";
        System.out.println("Mood:", mood);
    }
}