import {Component, Input} from '@angular/core';
import {TimeSlot} from "../../models/TimeSlot";

@Component({
  selector: 'app-time-slot',
  standalone: true,
  imports: [],
  templateUrl: './time-slot.component.html',
  styleUrl: './time-slot.component.scss'
})
export class TimeSlotComponent {

  @Input() timeSlot?: TimeSlot;

}
