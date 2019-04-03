import {AfterViewInit, Component, ElementRef, ViewChild} from '@angular/core';

@Component({
  selector: 'ngx-bit-editor',
  templateUrl: './ngx-bit-editor.component.html',
  styleUrls: ['./ngx-bit-editor.component.scss']
})
export class NgxBitEditorComponent implements AfterViewInit {

  @ViewChild('htmlTextAreaElement') htmlTextAreaElement: ElementRef;

  ngAfterViewInit() {
    console.log(this.htmlTextAreaElement.nativeElement);
  }
}
