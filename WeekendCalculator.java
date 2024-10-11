package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;
import java.util.stream.Stream;

public class WeekendCalculator {
    public static long countWeekendDays(LocalDate start, LocalDate end){

        if(start == null || end == null || start.isAfter(end)){
            throw new IllegalArgumentException("Invalid dates provided.");
        }
        long weekendCount = 0;
        LocalDate currentDate = start;

        while (!currentDate.isAfter(end)) {
            if (isWeekend(currentDate)) {
                weekendCount++;
            }
            currentDate = currentDate.plusDays(1); // Move to the next day
        }

        return weekendCount;
    } 
    
    private static boolean isWeekend(LocalDate date){
        DayOfWeek dayOfWeek = date.getDayOfWeek();
        return dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY;
    }
}
