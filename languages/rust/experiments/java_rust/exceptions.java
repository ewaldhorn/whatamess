class Exceptions {
    public static void main(String[] args) {
        var result = Exceptions.floorDivided(10, 4);
        System.out.println("10 divided by 4: " + result);

        try {
            var whoops = Exceptions.floorDivided(10, 0);
            System.out.println("10 divided by 0: " + whoops);
        } catch (Exception e) {
            System.out.println("Nope, let's not!");
        }

    }

    public static int floorDivided(float num, float by) {
        if (by == 0) {
            throw new RuntimeException("Division by zero.");
        }

        return (int) Math.floor(num /  by);
    }
}