package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;
import java.util.stream.Stream;

public class WeekendCalculator {
    public static long countWeekendDays(LocalDate start, LocalDate end){

        if(start == null || end == null || start.isAfter(end)){
            throw new IllegalArgumentException("Invalid dates provided");
        }
        return Stream.iterate(start, date -> date.plusDays(1))
                .limit(start.until(end.plusDays(1)).getDays())
                .filter(date-> isWeekend(date))
                .count();
    } 
    
    private static boolean isWeekend(LocalDate date){
        DayOfWeek dayOfWeek = date.getDayOfWeek();
        return dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY;
    }
}
