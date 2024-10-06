package sprint;

public class AreaCalculator {
    // Calculate the area of a square
    public double calculateArea(double side) {
        double area = side * side;
        return roundToTwoDecimals(area);
    }
    // Calculate the area of a rectangle
    public double calculateArea(double length, double width) {
        double area = length * width;
        return roundToTwoDecimals(area);
    }
     // Calculate the area of a circle, only if 'calculate' is true
     public double calculateArea(double radius, boolean calculate) {
        if (!calculate) {
            return Double.NaN;  // Return NaN if calculation is not allowed
        }
        double area = Math.PI * radius * radius;
        return roundToTwoDecimals(area);
    }
     // Helper method to round a value to 2 decimal places
     private double roundToTwoDecimals(double value) {
        return Math.round(value * 100.0) / 100.0;
    }
}
