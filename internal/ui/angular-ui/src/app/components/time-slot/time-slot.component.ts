import {Component, computed, input} from '@angular/core';
import {TimeSlot} from "../../models/TimeSlot";
import {DatePipe, NgStyle} from "@angular/common";
import {last} from "rxjs";
import {DurationPipe} from "../../duration.pipe";

@Component({
  selector: 'app-time-slot',
  standalone: true,
  imports: [
    NgStyle,
    DatePipe,
    DurationPipe
  ],
  templateUrl: './time-slot.component.html',
  styleUrl: './time-slot.component.scss'
})
export class TimeSlotComponent {

  timeSlot = input<TimeSlot>();
  scroll = input<number>(1);
  lastElement = input<boolean>(false);

  heightPx = computed(() => {
    const ts = this.timeSlot();
    if (!ts) return 55;
    const from = new Date(ts.Start);
    const to = new Date(ts.End);
    const durationInMinutes = (to.getTime() - from.getTime()) / (1000 * 60);
    let height = durationInMinutes * this.scroll()
    return Math.max(55, height);
  });

  protected readonly last = last;
}
