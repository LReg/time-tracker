import {Component, HostListener} from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {NavbarComponent} from "./components/navbar/navbar.component";
import {TimelineComponent} from "./components/timeline/timeline.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [NavbarComponent, TimelineComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'angular-ui';
  selectedDate: Date = new Date();



  handleDateSelected(date: Date) {
    this.selectedDate = date;
  }
}
