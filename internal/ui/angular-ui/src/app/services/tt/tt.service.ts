import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class TtService {

  private http = inject(HttpClient);

  quitTt() {
    this.http.get('http://localhost:8888/quit', {responseType: 'text'}).subscribe(() => {
      window.close();
    })
  }
}
