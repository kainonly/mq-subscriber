import {AfterViewInit, Component, ElementRef, forwardRef, OnInit, Renderer2, ViewChild} from '@angular/core';
import {NG_VALUE_ACCESSOR} from '@angular/forms';
import {DomSanitizer} from '@angular/platform-browser';

@Component({
  selector: 'ngx-bit-editor',
  templateUrl: './ngx-bit-editor.component.html',
  styleUrls: ['./ngx-bit-editor.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => NgxBitEditorComponent),
      multi: true,
    },
  ],
})
export class NgxBitEditorComponent implements OnInit, AfterViewInit {
  @ViewChild('htmlTextAreaElement') htmlTextAreaElement: ElementRef;

  private selfOnChange: (value: string) => void;
  private selfOnTouched: () => void;

  constructor(private renderer: Renderer2,
              private domSanitizer: DomSanitizer) {
  }

  writeValue(value: string) {
    this.renderer.setProperty(this.htmlTextAreaElement.nativeElement, 'value', value);
  }

  registerOnChange(fn: (_: any) => {}) {
    this.selfOnChange = fn;
  }

  registerOnTouched(fn: () => {}) {
    this.selfOnTouched = fn;
  }

  ngOnInit() {
  }

  ngAfterViewInit() {
  }

  inputText(event) {
    let HTMLString = event.target.value;
    HTMLString = HTMLString.replace(/\n/g, '<br/>');
    const safeHTML = this.domSanitizer.bypassSecurityTrustHtml(HTMLString);
    console.log(safeHTML);
  }
}
