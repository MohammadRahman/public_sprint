package sprint;

public class SmallestDivisor {
    public int smallestDivisor(int number) {
        // Implement logic here
        for (int i = 2; i <= number; i++) {
            if (number % i == 0) {
                return i;  // Return the first divisor found
            }
        }
        return number;
    }
}
