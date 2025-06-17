import {Component, computed, input, Input} from '@angular/core';
import {TimeSlot} from "../../models/TimeSlot";
import {DatePipe, JsonPipe, NgStyle} from "@angular/common";
import {last} from "rxjs";

@Component({
  selector: 'app-time-slot',
  standalone: true,
  imports: [
    NgStyle,
    DatePipe
  ],
  templateUrl: './time-slot.component.html',
  styleUrl: './time-slot.component.scss'
})
export class TimeSlotComponent {

  timeSlot = input<TimeSlot>();
  scroll = input<number>(1);
  @Input() lastElement = false;

  heightPx = computed(() => {
    const ts = this.timeSlot();
    if (!ts) return 40;
    const from = new Date(ts.Start);
    const to = new Date(ts.End);
    const durationInMinutes = (to.getTime() - from.getTime()) / (1000 * 60);
    return Math.max(40, durationInMinutes * this.scroll());
  });

  protected readonly last = last;
}
