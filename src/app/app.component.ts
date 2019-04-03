import {AfterViewInit, Component, ElementRef, ViewChild} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements AfterViewInit {
  @ViewChild('htmlTextAreaElement') htmlTextAreaElement: ElementRef;

  ngAfterViewInit() {
    console.log(this.htmlTextAreaElement.nativeElement);
  }
}
