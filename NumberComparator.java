package sprint;

public class NumberComparator {
    public String whichIsGreater(int n, double f) {
        if(n > f){
            return "Integer";
        }else if(n < f){
            return "Float";
        }else{
            return "Equal";
        }
    }
}
