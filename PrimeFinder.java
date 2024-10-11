package sprint;

import java.util.ArrayList;
import java.util.List;

public class PrimeFinder {
    public static List<Integer> findPrimesUpTo(int limit) {
        List<Integer> primes = new ArrayList<>();

        if (limit < 2) {
            return primes;
        }

        boolean[] isPrime = new boolean[limit + 1];

        for (int i = 2; i <= limit; i++) {
            isPrime[i] = true;
        }
        for (int p = 2; p * p <= limit; p++) {
            if (isPrime[p]) {
                // Mark all multiples of p as non-prime
                for (int multiple = p * p; multiple <= limit; multiple += p) {
                    isPrime[multiple] = false;
                }
            }
        }

        for (int i = 2; i <= limit; i++) {
            if (isPrime[i]) {
                primes.add(i);
            }
        }
        
        return primes;
    }
}
