public record User(long id, String email, int age) {
    // Compact constructor: input validation & normalization
    public User {
        Objects.requireNonNull(email, "email must not be null");
        if (age < 0) {
            throw new IllegalArgumentException("Age cannot be negative: " + age);
        }
        email = email.trim().toLowerCase(); // Normalize before assignment
    }
}

void main() {
  var user = new User(10, "some@onedo.com", 33);
  IO.println(user);
}
