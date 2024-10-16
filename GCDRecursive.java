package sprint;

public class GCDRecursive {
    public int gcd( int x, int y){
        if(x == 0 && y == 0){
            return 0;
        }
        if(y == 0){
            return Math.abs(x);
        }
        return gcd(y, x%y);
    }
}
