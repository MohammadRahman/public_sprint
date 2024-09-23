package sprint;

public class PrimeChecker {
    public static boolean isPrime(int number) {
        // Implement logic here
        if(number <= 1){
            return false;
        }
        for(int i = 2; i*i <= number; i++){
            if(number%i == 0){
                return false;
            }
        }
        return true;
    }
}
