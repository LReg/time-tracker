import {Component, EventEmitter, Output} from '@angular/core';
import {TtService} from "../../services/tt/tt.service";

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {

  @Output() selectedDate = new EventEmitter<Date>();

  constructor(
    private ttService: TtService
  ) {
  }

  handleChange($event: Event) {
      this.selectedDate.emit(($event.target as HTMLInputElement).valueAsDate || new Date());
  }

  handleQuit() {
    this.ttService.quitTt();
  }
}
