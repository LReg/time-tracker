import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {TimeSlot} from "../../models/TimeSlot";

@Injectable({
  providedIn: 'root'
})
export class TimeSlotService {

  constructor(private http: HttpClient) { }

  timeSlotsForDay$(date: Date) {
    this.http.get<TimeSlot[]>(`http://localhost:8888/tt/${date.getTime()}`);
  }

}
