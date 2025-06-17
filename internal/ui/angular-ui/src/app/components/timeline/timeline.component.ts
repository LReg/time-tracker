import {Component, HostListener, input} from '@angular/core';
import {toObservable} from "@angular/core/rxjs-interop";
import {TimeSlotService} from "../../services/TimeSlot/time-slot.service";
import {catchError, filter, of, switchMap} from "rxjs";
import {AsyncPipe} from "@angular/common";
import {TimeSlotComponent} from "../time-slot/time-slot.component";
import {TimeSlot} from "../../models/TimeSlot";

@Component({
  selector: 'app-timeline',
  standalone: true,
  imports: [
    AsyncPipe,
    TimeSlotComponent
  ],
  templateUrl: './timeline.component.html',
  styleUrl: './timeline.component.scss'
})
export class TimelineComponent {
  @HostListener('window:wheel', ['$event'])
  onWheel(event: WheelEvent) {
    if (event.deltaY > 0) {
      this.scroll -= .3;
    } else if (event.deltaY < 0) {
      this.scroll += .3;
    }
  }
  scroll = 1;
  date = input<Date>()
  timeSlots$ = toObservable(this.date).pipe(
    filter(date => date !== undefined),
    switchMap(date => this.timeSlotService.timeSlotsForDay$(date as Date)),
    catchError(() => of([] as TimeSlot[]))
  )
  constructor(
    private timeSlotService: TimeSlotService
  ) {
  }
}
