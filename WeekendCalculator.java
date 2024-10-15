package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;

public class WeekendCalculator {

    public static long countWeekendDays(LocalDate startDate, LocalDate endDate) {
        long weekendCount = 0;
        
        // Initialize the date variable before the loop
        LocalDate date = startDate;

        // Loop from the start date to the end date
        while (!date.isAfter(endDate)) {
            DayOfWeek dayOfWeek = date.getDayOfWeek();
            
            // Check if the day is Saturday or Sunday
            if (dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY) {
                weekendCount++;
            }

            // Move to the next day
            date = date.plusDays(1);
        }

        return weekendCount;
    }
}
