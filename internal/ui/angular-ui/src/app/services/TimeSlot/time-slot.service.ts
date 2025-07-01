import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {TimeSlot} from "../../models/TimeSlot";
import {catchError, of} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class TimeSlotService {

  private http = inject(HttpClient);

  timeSlotsForDay$(date: Date) {
    return this.http.get<TimeSlot[]>(`http://localhost:8888/tt/${date.getTime()}`).pipe(
        catchError(() => of([] as TimeSlot[]))
    );
  }

}
