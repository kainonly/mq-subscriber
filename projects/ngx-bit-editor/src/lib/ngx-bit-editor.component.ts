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

  static exec(id: string, value: string): boolean {
    return document.execCommand(id, false, value);
  }

  constructor(private renderer: Renderer2) {
  }

  writeValue(value: string) {
  }

  registerOnChange(fn: (_: any) => {}) {
    this.selfOnChange = fn;
  }

  registerOnTouched(fn: () => {}) {
    this.selfOnTouched = fn;
  }

  ngOnInit() {
    NgxBitEditorComponent.exec('defaultParagraphSeparator', 'p');
  }

  ngAfterViewInit() {
  }

  update() {
    // console.log(this.htmlDivElement);
    // const outerText = this.htmlDivElement.nativeElement.outerText;
    // const firstChild = this.htmlDivElement.nativeElement.firstChild;
    // if (firstChild.nodeType === 3) {
    //   // const firstNode = this.renderer.createElement('p');
    //   // this.renderer.setValue(firstChild, `<p>${firstChild.innerText}</p>`);
    // }
    console.log(this.htmlDivElement.nativeElement.innerHTML);
  }
}
