import {Component, HostListener, input} from '@angular/core';
import {toObservable} from "@angular/core/rxjs-interop";
import {TimeSlotService} from "../../services/TimeSlot/time-slot.service";
import {filter, switchMap} from "rxjs";
import {AsyncPipe} from "@angular/common";
import {TimeSlotComponent} from "../time-slot/time-slot.component";

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
      this.scroll -= 1;
    } else if (event.deltaY < 0) {
      this.scroll += 1;
    }
  }
  scroll = 1;
  date = input<Date>()
  timeSlots$ = toObservable(this.date).pipe(
    filter(date => date !== undefined),
    switchMap(date => this.timeSlotService.timeSlotsForDay$(date as Date))
  )
  constructor(
    private timeSlotService: TimeSlotService
  ) {
  }

  handleScroll($event: Event) {
    console.log($event)

  }
}
