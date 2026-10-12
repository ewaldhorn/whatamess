class Channel {
    protected final int timeoutMs;

    protected Channel(int timeoutMs) {
        this.timeoutMs = timeoutMs;
    }
}

public class SecureChannel extends Channel {
    private final String cipherKey;

    public SecureChannel(String rawKey, int timeoutMs) {
        // Java 25 (final): Statements before super(...)!
        if (rawKey == null || rawKey.length() < 32) {
            throw new IllegalArgumentException("Key too short or null");
        }
        int sanitizedTimeout = Math.clamp(timeoutMs, 100, 10_000);

        super(sanitizedTimeout);
        this.cipherKey = rawKey;
    }
}

void main() {
    var tmp = new SecureChannel("Bozo Ozo Froso Wozo Sozo Mozo Katozo Olozo", 0);
    IO.println(tmp.cipherKey);
}
