import {Component, EventEmitter, Output} from '@angular/core';
import {TtService} from "../../services/tt/tt.service";
import {FormsModule} from "@angular/forms";

@Component({
  selector: 'app-navbar',
  standalone: true,
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {

  now = new Date();
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
