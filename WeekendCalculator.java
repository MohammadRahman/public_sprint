package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;

public class WeekendCalculator {
    public static long countWeekendDays(LocalDate start, LocalDate end) {

        long weekendCount = 0;
        LocalDate currentDate = start;

        // if (start == null || end == null || start.isAfter(end)) {
        //     throw new IllegalArgumentException("Invalid dates provided.");
        // }

       

        while (!currentDate.isAfter(end)) {
            DayOfWeek dayOfWeek = date.getDayOfWeek();
            // if (isWeekend(currentDate)) {
            //     weekendCount++;
            // }
            // currentDate = currentDate.plusDays(1);
            if (dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY) {
                weekendCount++;
            }
            date = date.plusDays(1); 
        }

        return weekendCount;
    }

    // private static boolean isWeekend(LocalDate date) {
    //     DayOfWeek dayOfWeek = date.getDayOfWeek();
    //     return dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY;
    // }
}

