package sprint;

public class GCDCalculator {
    public int gcd(int a, int b) {
        // Implement logic here
        while (b != 0) {
            int tempV = b;
            b = a%b;
            a = tempV;
        }
        return Math.abs(a);
    }
}
