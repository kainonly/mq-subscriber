import {
  AfterViewInit,
  Component,
  ElementRef,
  forwardRef,
  OnInit,
  Renderer2,
  ViewChild,
} from '@angular/core';
import {NG_VALUE_ACCESSOR} from '@angular/forms';
import {NgxBitEditorService} from "./ngx-bit-editor.service";

@Component({
  selector: 'ngx-bit-editor',
  templateUrl: './ngx-bit-editor.component.html',
  styleUrls: ['./ngx-bit-editor.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => NgxBitEditorComponent),
      multi: true
    },
  ],
})
export class NgxBitEditorComponent implements OnInit, AfterViewInit {
  @ViewChild('htmlDivElement') htmlDivElement: ElementRef;

  html: string;

  private selfOnChange: (value: string) => void;
  private selfOnTouched: () => void;

  constructor(private renderer: Renderer2,
              private ngxBitEditorService: NgxBitEditorService) {
  }

  writeValue(value: string) {
    // console.log(value);
    // this.renderer.setProperty(this.htmlTextAreaElement.nativeElement, 'value', value);
  }

  registerOnChange(fn: (_: any) => {}) {
    this.selfOnChange = fn;
  }

  registerOnTouched(fn: () => {}) {
    this.selfOnTouched = fn;
  }

  ngOnInit() {
    this.ngxBitEditorService.exec('defaultParagraphSeparator', 'p');
  }

  ngAfterViewInit() {
  }

  updateHtml(event) {
    console.log(this.htmlDivElement.nativeElement.innerHTML);
  }
}
