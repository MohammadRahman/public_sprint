package sprint;

public class PowerCalculator {
    public int calculatePower(int base, int power) {
        // Implement logic here
        if(power == 0){
            return 1;
        }
        if(power < 0){
            throw new IllegalArgumentException("power must be non-negative.");
        }
        int result = 1;
        for(int i = 0; i < power; i++){
            result += base;
        }
        return result;
    }
}
